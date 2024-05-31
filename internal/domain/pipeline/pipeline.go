package pipeline

var _ Pipeline = (*pipeline)(nil)

func New(m Manifest) Pipeline {
	return &pipeline{m}
}

type pipeline struct {
	manifest Manifest
}

func (p *pipeline) Automated() bool {
	return p.manifest.Automated()
}

func (p *pipeline) Name() string {
	return p.manifest.Name()
}

func (p *pipeline) OutOfSync() (bool, error) {
}

func (p *pipeline) Sync(target Revision, c chan Status) {
	defer close(c)

	c <- StatusSyncing

	workingDir, err := getSrc(p.manifest.Remote())
	if err != nil {
		c <- StatusFailed
		return
	}

	a, err := build(p.manifest, workingDir, target)
	if err != nil {
		c <- StatusFailed
		return
	}

	// TODO: Stop running service

	s, err := install(p.manifest, a, target)
	if err != nil {
		c <- StatusFailed
		return
	}

	if err := s.Run(); err != nil {
		c <- StatusFailed
		return
	}

	c <- StatusSynced
}

func (p *pipeline) MarshalYAML() (interface{}, error) {
}
