package pipeline

import (
	"io"

	"github.com/go-yaml/yaml"
)

func NewManifestLoader() ManifestLoader {
	return &manifestLoader{}
}

type manifestLoader struct{}

func (*manifestLoader) Load(r io.Reader) (Manifest, error) {
	m := new(manifest)
	err := yaml.NewDecoder(r).Decode(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
