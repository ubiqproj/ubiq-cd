package pipeline

import (
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/service"
)

func install(m Manifest, a artifact.Artifact, r Revision) (service.Service, error) {
	err := m.registry().Register(a, m.Name(), artifact.Revision(r))
	if err != nil {
		return nil, err
	}
}
