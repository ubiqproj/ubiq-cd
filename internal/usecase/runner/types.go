package runner

import "ubiq-cd/internal/domain/pipeline"

type Runner interface {
	Apply(pipeline.Manifest) (pipeline.Pipeline, error)
	FindAll() ([]pipeline.Pipeline, error)
	Get(name string) (pipeline.Pipeline, error)
	Remove(pipeline.Pipeline) error

	Sync(p pipeline.Pipeline, to pipeline.Revision) error
}

type PipelineRepository interface {
	FindAll() ([]pipeline.Pipeline, error)
	Find(name string) (pipeline.Pipeline, error)
	Update(pipeline.Pipeline) error
	Remove(name string) error
}
