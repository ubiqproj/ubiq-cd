package docker

import (
	"ubiq-cd/internal/domain/artifact/docker/build"
	"ubiq-cd/internal/domain/artifact/docker/compose"
)

// Data of yaml manifest (`docker.service.[name]`)
type service struct {
	compose.Service
	BuildArgs *build.Args `yaml:"build,omitempty"`
}
