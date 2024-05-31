package docker

import (
	"fmt"
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/artifact/docker/build"
	"ubiq-cd/internal/domain/artifact/docker/compose"

	"github.com/go-yaml/yaml"
)

var _ artifact.Builder = (*Builder)(nil)

type Builder struct {
	data YamlStructure

	dockerPath string
}

type YamlStructure struct {
	Prepare  prepare
	Services map[string]service
	Volumes  map[string]compose.Volume
}

func (b *Builder) Build(
	workingDir artifact.AbsolutePath,
	group string,
	revision artifact.Revision,
) (artifact.Artifact, error) {
	err := runPrepareJob(workingDir, b.data.Prepare)
	if err != nil {
		return nil, err
	}

	services, err := buildServices(
		b.data.Services,
		workingDir, group, revision,
		build.WithDocker(b.dockerPath),
	)
	if err != nil {
		return nil, err
	}

	cf, err := compose.New(services, b.data.Volumes)
	if err != nil {
		return nil, err
	}

	return &Artifact{"compose.yaml", cf}, nil
}

func buildServices(
	services map[string]service,
	workingDir artifact.AbsolutePath,
	group string,
	revision artifact.Revision,
	options ...build.CmdEntryOptionApplier,
) (map[string]compose.Service, error) {
	return mapDict(services,
		func(name string, s service) (*compose.Service, error) {
			if /* build directive is not specified */ s.BuildArgs == nil {
				return nil, nil
			}
			tag := fmt.Sprintf("%s-%s:%s", group, name, revision)
			err := build.Run(workingDir, *s.BuildArgs, tag, options...)
			if err != nil {
				return nil, err
			}

			s.Service.Image = tag
			return &s.Service, nil
		},
	)
}

func mapDict[From, To any](v map[string]From, yield func(key string, s From) (*To, error)) (map[string]To, error) {
	res := make(map[string]To, len(v))
	for key, s1 := range v {
		s2, err := yield(key, s1)
		if err != nil {
			return nil, err
		}
		if s2 == nil {
			continue
		}
		res[key] = *s2
	}
	return res, nil
}

func (b *Builder) UnmarshalYAML(unmarshal func(interface{}) error) error {
	b.data = YamlStructure{}
	return unmarshal(&b.data)
}

func (b Builder) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(b.data)
}
