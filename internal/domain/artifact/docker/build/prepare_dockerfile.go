package build

import (
	"os"
	"path"
	"ubiq-cd/internal/domain/artifact"
	"ubiq-cd/internal/domain/artifact/docker/dockerfile"
)

func prepareDockerfile(contextDir artifact.AbsolutePath, args Args) (dockerfileName string, err error) {
	if /* dockerfile is specified */ args.DockerfileName != "" {
		return args.DockerfileName, nil
	}

	const FILENAME = "Dockerfile"
	f, err := os.Create(path.Join(string(contextDir), FILENAME))
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(dockerfile.New(args.Dockerfile))
	if err != nil {
		return "", err
	}

	return FILENAME, nil
}
