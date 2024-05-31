package pipeline

import "ubiq-cd/internal/domain/artifact"

func build(
	m Manifest,
	workingDir AbsolutePath,
	r Revision,
) (artifact.Artifact, error) {
	return m.builder().Build(artifact.AbsolutePath(workingDir), string(m.Name()), r)
}
