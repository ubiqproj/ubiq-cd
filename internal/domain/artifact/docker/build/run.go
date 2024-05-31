package build

import (
	"path"
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/artifact/docker/dockerfile"
)

// Data of yaml manifest (`docker.service.[name].build`)
type Args struct {
	Context        string             `yaml:"context,omitempty"`
	BuildArgs      map[string]string  `yaml:"args,omitempty"`
	DockerfileName string             `yaml:"dockerfile,omitempty"`
	Dockerfile     []dockerfile.Stage `yaml:"stages,omitempty"`
}

func Run(workingDir artifact.AbsolutePath, args Args, tag string, options ...CmdEntryOptionApplier) error {
	absoluteContextWorkingDir := artifact.AbsolutePath(
		path.Join(string(workingDir), args.Context),
	)
	dockerfileName, err := prepareDockerfile(absoluteContextWorkingDir, args)
	if err != nil {
		return err
	}

	options = append(options,
		WithTag(tag),
		WithBuildArgs(args.BuildArgs),
		withDockerfile(dockerfileName),
	)
	return execCmd(absoluteContextWorkingDir, options...)
}
