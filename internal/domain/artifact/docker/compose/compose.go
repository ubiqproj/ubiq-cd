package compose

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

type Service struct {
	Image        string
	Ports        []string          `yaml:"ports,omitempty"`
	DependsOn    []string          `yaml:"depends_on,omitempty"`
	VolumeMounts []VolumeMount     `yaml:"volumes,omitempty"`
	Envinronment map[string]string `yaml:"environment,omitempty"`
}

type VolumeMount struct {
	VolumeName string `yaml:"volume"`
	TargetPath string `yaml:"target"`
}

func (vm VolumeMount) MarshalYAML() (interface{}, error) {
	return fmt.Sprintf("%s:%s", vm.VolumeName, vm.TargetPath), nil
}

type Volume struct {
	External bool `yaml:"external,omitempty"`
}

type YamlStructure struct {
	Version  string
	Services map[string]Service
	Volumes  map[string]Volume `yaml:"volumes,omitempty"`
}

func New(services map[string]Service, volumes map[string]Volume) ([]byte, error) {
	// TODO: Support network directive
	return yaml.Marshal(YamlStructure{"3", services, volumes})
}
