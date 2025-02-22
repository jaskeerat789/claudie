package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/Berops/claudie/internal/envs"
	"github.com/Berops/claudie/internal/healthcheck"
	"github.com/Berops/claudie/internal/kubectl"
	"github.com/Berops/claudie/internal/utils"
	"github.com/Berops/claudie/proto/pb"
	"github.com/Berops/claudie/services/kuber/server/longhorn"
	"github.com/Berops/claudie/services/kuber/server/nodes"
	"github.com/Berops/claudie/services/kuber/server/secret"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	defaultKuberPort = 50057
	outputDir        = "services/kuber/server/clusters"
)

type (
	IPPair struct {
		PublicIP  net.IP `json:"public_ip"`
		PrivateIP net.IP `json:"private_ip"`
	}

	ClusterMetadata struct {
		// NodeIps maps node-name to public-private ip pairs.
		NodeIps map[string]IPPair `json:"node_ips"`
		// PrivateKey is the private SSH key for the nodes.
		PrivateKey string `json:"private_key"`
	}
)

type server struct {
	pb.UnimplementedKuberServiceServer
}

func (s *server) SetUpStorage(ctx context.Context, req *pb.SetUpStorageRequest) (*pb.SetUpStorageResponse, error) {
	desiredState := req.GetDesiredState()
	var errGroup errgroup.Group
	for _, cluster := range desiredState.GetClusters() {
		func(c *pb.K8Scluster) {
			errGroup.Go(func() error {
				clusterID := fmt.Sprintf("%s-%s", c.ClusterInfo.Name, c.ClusterInfo.Hash)
				clusterDir := filepath.Join(outputDir, clusterID)
				longhorn := longhorn.Longhorn{Cluster: c, Directory: clusterDir}
				err := longhorn.SetUp()
				if err != nil {
					return fmt.Errorf("error while setting up the longhorn for %s : %w", clusterID, err)
				}
				log.Info().Msgf("Longhorn successfully set up on the cluster %s", clusterID)
				return nil
			})
		}(cluster)
	}
	if err := errGroup.Wait(); err != nil {
		log.Error().Msgf("Error encountered in SetUpStorage : %s", err.Error())
		return &pb.SetUpStorageResponse{DesiredState: desiredState}, err
	}
	log.Info().Msgf("Storage was successfully set up for project %s", desiredState.Name)
	return &pb.SetUpStorageResponse{DesiredState: desiredState}, nil
}

func (s *server) StoreClusterMetadata(ctx context.Context, req *pb.StoreClusterMetadataRequest) (*pb.StoreClusterMetadataResponse, error) {
	md := ClusterMetadata{
		NodeIps:    make(map[string]IPPair),
		PrivateKey: req.GetCluster().GetClusterInfo().GetPrivateKey(),
	}

	for _, pool := range req.GetCluster().GetClusterInfo().GetNodePools() {
		for _, node := range pool.GetNodes() {
			md.NodeIps[node.Name] = IPPair{
				PublicIP:  net.ParseIP(node.Public),
				PrivateIP: net.ParseIP(node.Private),
			}
		}
	}

	b, err := json.Marshal(md)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s cluster metadata: %w", req.GetCluster().GetClusterInfo().GetName(), err)
	}

	// local deployment - print metadata
	if namespace := envs.Namespace; namespace == "" {
		// NOTE: DEBUG print
		// var buffer bytes.Buffer
		// for node, ips := range md.NodeIps {
		// 	buffer.WriteString(fmt.Sprintf("%s: %v \t| %v \n", node, ips.PublicIP, ips.PrivateIP))
		// }
		// buffer.WriteString(fmt.Sprintf("%s\n", md.PrivateKey))
		// log.Info().Msgf("Cluster metadata from cluster %s \n%s", req.GetCluster().ClusterInfo.Name, buffer.String())
		return &pb.StoreClusterMetadataResponse{}, nil
	}

	clusterID := fmt.Sprintf("%s-%s", req.GetCluster().ClusterInfo.Name, req.GetCluster().ClusterInfo.Hash)
	clusterDir := filepath.Join(outputDir, clusterID)
	sec := secret.New(clusterDir, secret.NewYaml(
		secret.Metadata{Name: fmt.Sprintf("%s-metadata", clusterID)},
		secret.Data{SecretData: base64.StdEncoding.EncodeToString(b)},
	))

	if err := sec.Apply(envs.Namespace, ""); err != nil {
		log.Error().Msgf("Failed to store cluster metadata for %s: %s", req.Cluster.ClusterInfo.Name, err)
		return nil, fmt.Errorf("error while creating cluster metadata secret for %s", req.Cluster.ClusterInfo.Name)
	}

	log.Info().Msgf("Cluster Metadata was successfully stored for cluster %s", req.Cluster.ClusterInfo.Name)
	return &pb.StoreClusterMetadataResponse{}, nil
}

func (s *server) DeleteClusterMetadata(ctx context.Context, req *pb.DeleteClusterMetadataRequest) (*pb.DeleteClusterMetadataResponse, error) {
	namespace := envs.Namespace
	if namespace == "" {
		return &pb.DeleteClusterMetadataResponse{}, nil
	}

	kc := kubectl.Kubectl{}
	secretName := fmt.Sprintf("%s-%s-metadata", req.Cluster.ClusterInfo.Name, req.Cluster.ClusterInfo.Hash)
	if err := kc.KubectlDeleteResource("secret", secretName, namespace); err != nil {
		log.Error().Msgf("Failed to remove cluster metadata for %s: %s", req.Cluster.ClusterInfo.Name, err)
		return nil, fmt.Errorf("error while deleting kubeconfig secret for %s: %w", req.Cluster.ClusterInfo.Name, err)
	}

	log.Info().Msgf("Deleted ClusterMetadata secret for cluster %s", req.Cluster.ClusterInfo.Name)
	return &pb.DeleteClusterMetadataResponse{}, nil
}

func (s *server) StoreKubeconfig(ctx context.Context, req *pb.StoreKubeconfigRequest) (*pb.StoreKubeconfigResponse, error) {
	// local deployment - print kubeconfig
	if namespace := envs.Namespace; namespace == "" {
		//NOTE: DEBUG print
		// log.Info().Msgf("The kubeconfig for %s\n%s:", clusterID,cluster.Kubeconfig)
		return &pb.StoreKubeconfigResponse{}, nil
	}
	cluster := req.GetCluster()
	clusterID := fmt.Sprintf("%s-%s", cluster.ClusterInfo.Name, cluster.ClusterInfo.Hash)

	clusterDir := filepath.Join(outputDir, clusterID)
	sec := secret.New(clusterDir, secret.NewYaml(
		secret.Metadata{Name: fmt.Sprintf("%s-kubeconfig", clusterID)},
		secret.Data{SecretData: base64.StdEncoding.EncodeToString([]byte(cluster.GetKubeconfig()))},
	))

	if err := sec.Apply(envs.Namespace, ""); err != nil {
		log.Error().Msgf("Failed to store kubeconfig for %s: %s", cluster.ClusterInfo.Name, err)
		return nil, fmt.Errorf("error while creating the kubeconfig secret for %s", cluster.ClusterInfo.Name)
	}

	log.Info().Msgf("Kubeconfig was successfully stored for cluster %s", cluster.ClusterInfo.Name)
	return &pb.StoreKubeconfigResponse{}, nil
}

func (s *server) DeleteKubeconfig(ctx context.Context, req *pb.DeleteKubeconfigRequest) (*pb.DeleteKubeconfigResponse, error) {
	namespace := envs.Namespace
	if namespace == "" {
		return &pb.DeleteKubeconfigResponse{}, nil
	}

	kc := kubectl.Kubectl{}
	cluster := req.Cluster
	secretName := fmt.Sprintf("%s-%s-kubeconfig", cluster.ClusterInfo.Name, cluster.ClusterInfo.Hash)

	if err := kc.KubectlDeleteResource("secret", secretName, namespace); err != nil {
		log.Error().Msgf("Failed to remove kubeconfig for %s: %s", cluster.ClusterInfo.Name, err)
		return nil, fmt.Errorf("error while deleting kubeconfig secret for %s : %w", cluster.ClusterInfo.Name, err)
	}

	log.Info().Msgf("Deleted kubeconfig secret for cluster %s", cluster.ClusterInfo.Name)
	return &pb.DeleteKubeconfigResponse{}, nil
}

func (s *server) DeleteNodes(ctx context.Context, req *pb.DeleteNodesRequest) (*pb.DeleteNodesResponse, error) {
	deleter := nodes.New(req.MasterNodes, req.WorkerNodes, req.Cluster)
	cluster, err := deleter.DeleteNodes()
	if err != nil {
		log.Error().Msgf("Error while deleting nodes for %s : %s", req.Cluster.ClusterInfo.Name, err.Error())
		return &pb.DeleteNodesResponse{}, err
	}
	log.Info().Msgf("Nodes for cluster %s were successfully deleted", req.Cluster.ClusterInfo.Name)
	return &pb.DeleteNodesResponse{Cluster: cluster}, nil
}

func main() {
	// initialize logger
	utils.InitLog("kuber")

	// Set the kuber port
	kuberPort := utils.GetenvOr("KUBER_PORT", fmt.Sprint(defaultKuberPort))

	// Start Terraformer Service
	trfAddr := net.JoinHostPort("0.0.0.0", kuberPort)
	lis, err := net.Listen("tcp", trfAddr)
	if err != nil {
		log.Fatal().Msgf("Failed to listen on %v", err)
	}
	log.Info().Msgf("Kuber service is listening on: %s", trfAddr)

	s := grpc.NewServer()
	pb.RegisterKuberServiceServer(s, &server{})

	// Add health service to gRPC
	healthService := healthcheck.NewServerHealthChecker(kuberPort, "KUBER_PORT", nil)
	grpc_health_v1.RegisterHealthServer(s, healthService)

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		defer signal.Stop(ch)

		// wait for either the received signal or
		// check if an error occurred in other
		// go-routines.
		var err error
		select {
		case <-ctx.Done():
			err = ctx.Err()
		case sig := <-ch:
			log.Info().Msgf("Received signal %v", sig)
			err = errors.New("kuber interrupt signal")
		}

		log.Info().Msg("Gracefully shutting down gRPC server")
		s.GracefulStop()

		// Sometimes when the container terminates gRPC logs the following message:
		// rpc error: code = Unknown desc = Error: No such container: hash of the container...
		// It does not affect anything as everything will get terminated gracefully
		// this time.Sleep fixes it so that the message won't be logged.
		time.Sleep(1 * time.Second)

		return err
	})

	g.Go(func() error {
		// s.Serve() will create a service goroutine for each connection
		if err := s.Serve(lis); err != nil {
			return fmt.Errorf("kuber failed to serve: %w", err)
		}
		log.Info().Msg("Finished listening for incoming connections")
		return nil
	})

	log.Info().Msgf("Stopping Kuber: %v", g.Wait())
}
