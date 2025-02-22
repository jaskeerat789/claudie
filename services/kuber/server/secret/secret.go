package secret

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Berops/claudie/internal/kubectl"
	"gopkg.in/yaml.v3"
)

// Secret holds information necessary to create a secret
type Secret struct {
	// Directory - directory where secret will be created
	Directory string
	// YamlManifest - secret specification
	YamlManifest SecretYaml
}

type SecretYaml struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	SecretType string   `yaml:"type"`
	Data       Data     `yaml:"data"`
}

type Metadata struct {
	Name   string      `yaml:"name"`
	Labels interface{} `yaml:"labels"`
}

type Data struct {
	SecretData string
}

const (
	filePermission os.FileMode = 0644
	filename                   = "secret.yaml"
)

// New create a k8s Secret manifest object from the specified manifest.
func New(directory string, secretYaml SecretYaml) Secret {
	return Secret{
		Directory:    directory,
		YamlManifest: secretYaml,
	}
}

// NewYaml created a template with pre-defined defaults and optional metadata & data fields.
func NewYaml(md Metadata, data Data) SecretYaml {
	return SecretYaml{
		APIVersion: "v1",
		Kind:       "Secret",
		Metadata:   md,
		SecretType: "Opaque",
		Data:       data,
	}
}

// Apply creates a secret manifests and applies it in the cluster (specified by given kubeconfig) in the specified namespace
// if the kubeconfig is left empty, it uses default kubeconfig
func (s *Secret) Apply(namespace, kubeconfig string) error {
	// setting empty string for kubeconfig will create secret on same cluster where claudie is running
	kubectl := kubectl.Kubectl{Kubeconfig: kubeconfig}
	path := filepath.Join(s.Directory, filename)

	if _, err := os.Stat(s.Directory); os.IsNotExist(err) {
		if err := os.Mkdir(s.Directory, os.ModePerm); err != nil {
			return fmt.Errorf("could not create a directory for %s: %w", s.YamlManifest.Metadata.Name, err)
		}
	}

	if err := s.saveSecretManifest(path); err != nil {
		return fmt.Errorf("error while saving secret.yaml for %s : %w", s.YamlManifest.Metadata.Name, err)
	}

	if err := kubectl.KubectlApply(path, namespace); err != nil {
		return fmt.Errorf("error while applying secret.yaml for %s : %w", s.YamlManifest.Metadata.Name, err)
	}

	// clean up
	if err := os.RemoveAll(s.Directory); err != nil {
		return fmt.Errorf("error while delete the secret.yaml for %s : %w", s.YamlManifest.Metadata.Name, err)
	}
	return nil
}

// saves secret into the file system
func (s *Secret) saveSecretManifest(path string) error {
	secretYaml, err := yaml.Marshal(&s.YamlManifest)
	if err != nil {
		return fmt.Errorf("failed to marshal secret manifest yaml for %s : %w", path, err)
	}

	if err = os.WriteFile(path, secretYaml, filePermission); err != nil {
		return fmt.Errorf("error while saving secret manifest file %s : %w", path, err)
	}
	return nil
}
