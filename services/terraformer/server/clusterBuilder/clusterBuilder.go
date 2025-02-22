package clusterBuilder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	comm "github.com/Berops/claudie/internal/command"
	"github.com/Berops/claudie/internal/templateUtils"
	"github.com/Berops/claudie/internal/utils"
	"github.com/Berops/claudie/proto/pb"
	"github.com/Berops/claudie/services/terraformer/server/backend"
	"github.com/Berops/claudie/services/terraformer/server/provider"
	"github.com/Berops/claudie/services/terraformer/server/terraform"
	"github.com/rs/zerolog/log"
)

const Output = "services/terraformer/server/clusters"

// ClusterBuilder wraps data needed for building a cluster.
type ClusterBuilder struct {
	// DesiredInfo contains the information about the
	// desired state of the cluster.
	DesiredInfo *pb.ClusterInfo
	// CurrentInfo contains the information about the
	// current state of the cluster.
	CurrentInfo *pb.ClusterInfo
	// ProjectName is the name of the manifest.
	ProjectName string
	// ClusterType is the type of the cluster being build
	// LoadBalancer or K8s.
	ClusterType pb.ClusterType
	// Metadata contains data that further describe
	// the cluster that is to be build. For example,
	// in the case of LoadBalancer this will contain the defined
	// roles from the manifest. Can be nil if no data is supplied.
	Metadata map[string]any
}

type NodepoolsData struct {
	ClusterName string
	ClusterHash string
	NodePools   []*pb.NodePool
	Metadata    map[string]any
	Regions     []string
}

type outputNodepools struct {
	IPs map[string]interface{} `json:"-"`
}

func (c ClusterBuilder) CreateNodepools() error {
	clusterID := fmt.Sprintf("%s-%s", c.DesiredInfo.Name, c.DesiredInfo.Hash)
	clusterDir := filepath.Join(Output, clusterID)

	if err := c.generateFiles(clusterID, clusterDir); err != nil {
		return fmt.Errorf("failed to generate files: %w", err)
	}

	terraform := terraform.Terraform{
		Directory: clusterDir,
		StdOut:    comm.GetStdOut(clusterID),
		StdErr:    comm.GetStdErr(clusterID),
	}

	if err := terraform.TerraformInit(); err != nil {
		return fmt.Errorf("error while running terraform init in %s : %w", clusterID, err)
	}

	if err := terraform.TerraformApply(); err != nil {
		return fmt.Errorf("error while running terraform apply in %s : %w", clusterID, err)
	}
	oldNodes := c.getCurrentNodes()

	// fill new nodes with output
	for _, nodepool := range c.DesiredInfo.NodePools {
		output, err := terraform.TerraformOutput(nodepool.Name)
		if err != nil {
			return fmt.Errorf("error while getting output from terraform for %s : %w", nodepool.Name, err)
		}
		out, err := readIPs(output)
		if err != nil {
			return fmt.Errorf("error while reading the terraform output for %s : %w", nodepool.Name, err)
		}
		fillNodes(&out, nodepool, oldNodes)
	}

	// Clean after terraform
	if err := os.RemoveAll(clusterDir); err != nil {
		return fmt.Errorf("error while deleting files in %s : %w", clusterDir, err)
	}

	return nil
}

func (c ClusterBuilder) DestroyNodepools() error {
	clusterID := fmt.Sprintf("%s-%s", c.CurrentInfo.Name, c.CurrentInfo.Hash)
	clusterDir := filepath.Join(Output, clusterID)

	if err := c.generateFiles(clusterID, clusterDir); err != nil {
		return fmt.Errorf("failed to generate files: %w", err)
	}

	terraform := terraform.Terraform{
		Directory: clusterDir,
		StdOut:    comm.GetStdOut(clusterID),
		StdErr:    comm.GetStdErr(clusterID),
	}

	if err := terraform.TerraformInit(); err != nil {
		return fmt.Errorf("error while running terraform init in %s : %w", clusterID, err)
	}

	if err := terraform.TerraformDestroy(); err != nil {
		return fmt.Errorf("error while running terraform apply in %s : %w", clusterID, err)
	}

	// Clean after terraform.
	if err := os.RemoveAll(clusterDir); err != nil {
		return fmt.Errorf("error while deleting files in %s : %w", clusterDir, err)
	}

	return nil
}

func (c ClusterBuilder) generateFiles(clusterID, clusterDir string) error {
	// generate backend
	backend := backend.Backend{ProjectName: c.ProjectName, ClusterName: clusterID, Directory: clusterDir}
	err := backend.CreateFiles()
	if err != nil {
		return err
	}

	// generate .tf files for nodepools
	var clusterInfo *pb.ClusterInfo
	template := templateUtils.Templates{Directory: clusterDir}
	templateLoader := templateUtils.TemplateLoader{Directory: templateUtils.TerraformerTemplates}

	if c.DesiredInfo != nil {
		clusterInfo = c.DesiredInfo
	} else if c.CurrentInfo != nil {
		clusterInfo = c.CurrentInfo
	}

	// generate Providers terraform configuration
	providers := provider.Provider{ProjectName: c.ProjectName, ClusterName: clusterID, Directory: clusterDir}
	err = providers.CreateProvider(clusterInfo)
	if err != nil {
		return err
	}

	tplType := getTplFile(c.ClusterType)
	//sort nodepools by a provider
	sortedNodePools := utils.GroupNodepoolsByProviderSpecName(clusterInfo)
	for providerSpecName, nodepools := range sortedNodePools {
		// list all regions being used for this provider
		regions := utils.GetRegions(nodepools)

		// based on the cluster type fill out the nodepools data to be used
		nodepoolData := NodepoolsData{
			NodePools:   nodepools,
			ClusterName: clusterInfo.Name,
			ClusterHash: clusterInfo.Hash,
			Metadata:    c.Metadata,
			Regions:     regions,
		}

		// Load TF files of the specific cloud provider
		tpl, err := templateLoader.LoadTemplate(fmt.Sprintf("%s%s", nodepools[0].Provider.CloudProviderName, tplType))
		if err != nil {
			return fmt.Errorf("error while parsing template file %s : %w", fmt.Sprintf("%s%s", nodepools[0].Provider.CloudProviderName, tplType), err)
		}

		// Parse the templates and create Tf files
		err = template.Generate(tpl, fmt.Sprintf("%s-%s.tf", clusterID, providerSpecName), nodepoolData)
		if err != nil {
			return fmt.Errorf("error while generating %s file : %w", fmt.Sprintf("%s-%s.tf", clusterID, providerSpecName), err)
		}

		// Create publicKey file for a cluster
		if err := utils.CreateKeyFile(clusterInfo.PublicKey, clusterDir, "public.pem"); err != nil {
			return fmt.Errorf("error creating key file for %s : %w", clusterDir, err)
		}

		// save keys
		if err = utils.CreateKeyFile(nodepools[0].Provider.Credentials, clusterDir, providerSpecName); err != nil {
			return fmt.Errorf("error creating provider credential key file for provider %s in %s : %w", providerSpecName, clusterDir, err)
		}
	}

	return nil
}

// getCurrentNodes returns all nodes which are in a current state
func (c *ClusterBuilder) getCurrentNodes() []*pb.Node {
	// group all the nodes together to make searching with respect to IP easy
	var oldNodes []*pb.Node
	if c.CurrentInfo != nil {
		for _, oldNodepool := range c.CurrentInfo.NodePools {
			oldNodes = append(oldNodes, oldNodepool.Nodes...)
		}
	}
	return oldNodes
}

// fillNodes creates pb.Node slices in desired state, with the new nodes and old nodes
func fillNodes(terraformOutput *outputNodepools, newNodePool *pb.NodePool, oldNodes []*pb.Node) {
	// fill slices from terraformOutput maps with names of nodes to ensure an order
	var tempNodes []*pb.Node
	// get sorted list of keys
	sortedNodeNames := getKeysFromMap(terraformOutput.IPs)
	for _, nodeName := range sortedNodeNames {
		var nodeType pb.NodeType
		var private string

		if newNodePool.IsControl {
			nodeType = pb.NodeType_master
		} else {
			nodeType = pb.NodeType_worker
		}
		if len(oldNodes) > 0 {
			for _, node := range oldNodes {
				//check if node was defined before
				if fmt.Sprint(terraformOutput.IPs[nodeName]) == node.Public && nodeName == node.Name {
					log.Debug().Msgf("Carrying information to desired state for node %s", nodeName)
					// carry privateIP to desired state, so it will not get overwritten in Ansibler
					private = node.Private
					// carry node type since it might be API endpoint, which should not change once set
					nodeType = node.NodeType
					break
				}
			}
		}

		tempNodes = append(tempNodes, &pb.Node{
			Name:     nodeName,
			Public:   fmt.Sprint(terraformOutput.IPs[nodeName]),
			Private:  private,
			NodeType: nodeType,
		})
	}
	newNodePool.Nodes = tempNodes
}

// getKeysFromMap returns an array of all keys in a map
func getKeysFromMap(data map[string]interface{}) []string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// readIPs reads json output format from terraform and unmarshal it into map[string]map[string]string readable by GO
func readIPs(data string) (outputNodepools, error) {
	var result outputNodepools
	// Unmarshal or Decode the JSON to the interface.
	err := json.Unmarshal([]byte(data), &result.IPs)
	return result, err
}

func getTplFile(clusterType pb.ClusterType) string {
	switch clusterType {
	case pb.ClusterType_K8s:
		return ".tpl"
	case pb.ClusterType_LB:
		return "-lb.tpl"
	}
	return ""
}
