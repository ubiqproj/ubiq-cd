package pipeline

import (
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/artifact/docker"
)

var _ Manifest = (*manifest)(nil)

type manifest struct {
	name      string
	remote    Remote
	automated bool

	_builder  artifact.Builder
	_registry artifact.Registry
}

// Remote implements Manifest.
func (m *manifest) Remote() Remote {
	return m.remote
}

// Name implements Manifest.
func (m *manifest) Name() string {
	return m.name
}

// Automated implements Manifest.
func (m *manifest) Automated() bool {
	return m.automated
}

// builder implements Manifest.
func (m *manifest) builder() artifact.Builder {
	return m._builder
}

// registry implements Manifest.
func (m *manifest) registry() artifact.Registry {
	return m._registry
}

// UnmarshalYAML implements Manifest.
func (m *manifest) UnmarshalYAML(unmarshal func(interface{}) error) error {
	r := struct {
		Name      string
		Remote    string `yaml:"git_remote_url"`
		Automated bool   `yaml:"auto_sync"`
		Docker    docker.Builder
	}{}
	err := unmarshal(&r)
	if err != nil {
		return err
	}

	m.name = r.Name
	m.remote = Remote(r.Remote)
	m.automated = r.Automated
	m._builder = &r.Docker
}
