package runner

import "ubiq-cd/internal/domain/pipeline"

var _ Runner = (*runner)(nil)

type runner struct {
	repository PipelineRepository
}

// Apply implements Runner.
func (r *runner) Apply(m pipeline.Manifest) (pipeline.Pipeline, error) {
	return pipeline.New(m), nil
}

// Find implements Runner.
func (r *runner) FindAll() ([]pipeline.Pipeline, error) {
	return r.repository.FindAll()
}

// Get implements Runner.
func (r *runner) Get(name string) (pipeline.Pipeline, error) {
	return r.repository.Find(name)
}

// Remove implements Runner.
func (r *runner) Remove(p pipeline.Pipeline) error {
	return r.repository.Remove(p.Name())
}

// Sync implements Runner.
func (r *runner) Sync(p pipeline.Pipeline, to pipeline.Revision) error {
	c := make(chan pipeline.Status, 2 /* Syncing -> Failed or Synced */)
	go p.Sync(to, c)
	for status := range c {
		// Update pipeline status of mutex...
		switch status {
		case pipeline.StatusSyncing:
			// ...
		case pipeline.StatusFailed:
			// ...
		case pipeline.StatusSynced:
			// ...
		default:
			// ...
		}
	}

	// ...

}
