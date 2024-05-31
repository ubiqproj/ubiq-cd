package pipeline

import (
	"io"
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/service"

	"github.com/go-yaml/yaml"
)

type Pipeline interface {
	Name() string
	OutOfSync() (bool, error)
	Automated() bool
	Sync(to Revision, c chan Status)

	yaml.Marshaler
}

type Status string

var (
	StatusSynced    Status = "synced"
	StatusSyncing   Status = "syncing"
	StatusOutOfSync Status = "out of sync"
	StatusFailed    Status = "failed"
)

type PipelineState interface {
	Status() Status
	Revision() Revision
	Artifact() artifact.Artifact
}

type Job interface {
	GetSrc() error
	Build() (artifact.Revision, error)
	Install() (service.Service, error)
}

type Manifest interface {
	Name() string
	Remote() Remote
	Automated() bool
	builder() artifact.Builder
	registry() artifact.Registry

	yaml.Unmarshaler
}

type ManifestLoader interface {
	Load(r io.Reader) (Manifest, error)
}

// Git
type (
	Remote   string
	Branch   string
	Revision string
	Regex    string

	AbsolutePath string

	Git interface {
		CloneRecursively(pipelineName string, r Remote, b Branch) (AbsolutePath, error)
		CheckUpdate(pipelineName string, r Remote, current Revision, option interface{ apply(*GitCheckUpdateOption) }) (bool, error)
	}

	GitCheckUpdateOption struct {
		GitCheckUpdateOptionBranch   *Branch
		GitCheckUpdateOptionTagRegex *Regex
	}
)
