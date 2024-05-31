package artifact

import "github.com/go-yaml/yaml"

// Revision format is
//
//	"<branch>-<first 8 characters of revision>-<first 8 characters of job id>" or
//	"<tag>-<first 8 characters of job id>"
//
// e.g., "main-13ab4d53-e9f455bf", "v1.0.0-e9f455bf"
type Revision string

type AbsolutePath string

type Artifact interface {
	Save(dir AbsolutePath) error
}

type Builder interface {
	Build(
		workingDir AbsolutePath,
		group string,
		r Revision,
	) (Artifact, error)

	yaml.Unmarshaler
	yaml.Marshaler
}

type Registry interface {
	Register(a Artifact, group string, r Revision) error
	Exists(group string, r Revision) (bool, error)
	Load() (Artifact, error)
}

type Runner interface {
	Run(Artifact) error
	Stop(Artifact) error
}
