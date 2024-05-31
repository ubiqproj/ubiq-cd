package docker

import (
	"os"
	"path"
	"ubiq-cd/internal/domain/artifact"
)

type prepare struct {
	WriteFiles []writeFile `yaml:"write_files"`
	// TODO
	// Jobs       []struct {
	// 	Name      string
	// 	Image     string
	// 	Args      map[string]string
	// 	Commands  []string `yaml:"run"`
	// 	Artifacts []struct {
	// 		Path   string
	// 		SaveTo string `yaml:"to"`
	// 	}
	// }
}

func runPrepareJob(workingDir artifact.AbsolutePath, p prepare) error {
	return writeFiles(workingDir, p.WriteFiles...)
}

type writeFile struct {
	Path    string
	Content string
}

func writeFiles(workingDir artifact.AbsolutePath, files ...writeFile) error {
	for _, file := range files {
		f, err := os.Create(path.Join(string(workingDir), file.Path))
		if err != nil {
			return err
		}

		_, err = f.WriteString(file.Content)
		if err != nil {
			return err
		}

		f.Close()
	}
	return nil
}
