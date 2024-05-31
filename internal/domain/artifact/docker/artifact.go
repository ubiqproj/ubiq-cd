package docker

import (
	"os"
	"path"
	"ubiq-cd/internal/domain/artifact"
)

var _ artifact.Artifact = (*Artifact)(nil)

type Artifact struct {
	filename string
	data     []byte /* data of compose.yaml */
}

func (a *Artifact) Save(dir artifact.AbsolutePath) error {
	f, err := os.Create(path.Join(string(dir), a.filename))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(a.data)
	if err != nil {
		return err
	}

	return nil
}
