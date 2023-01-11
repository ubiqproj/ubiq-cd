package systemd

import (
	"bytes"
	"fmt"
	"strings"
	"systemd-cd/domain/model/logger"
	"systemd-cd/domain/model/toml"
)

// writeEnvFile implements iSystemdService
func (s Systemd) writeEnvFile(e map[string]string, path string) error {
	logger.Logger().Tracef("Called:\n\targ.e: %v\n\targ.path %v", e, path)

	// Check env file path and mkdir
	if strings.HasSuffix(path, "/") {
		err := fmt.Errorf("service env file path %v is not a file", path)
		logger.Logger().Errorf("Error:\n\terr: %v", err)
		return err
	}
	if !strings.HasPrefix(path, "/") {
		err := fmt.Errorf("service env file path %v must be absolute", path)
		logger.Logger().Errorf("Error:\n\terr: %v", err)
		return err
	}
	err := mkdirIfNotExist("/" + strings.Join(strings.Split(path, "/")[1:len(strings.Split(path, "/"))-1], "/"))
	if err != nil {
		logger.Logger().Errorf("Error:\n\terr: %v", err)
		return err
	}

	// Encode
	b := &bytes.Buffer{}
	// Add annotation
	b.WriteString("#! Generated by systemd-cd\n")
	indent := ""
	// TODO: toml format で問題ないか検証
	logger.Logger().Warn("Unchecked code: no problem to systemd service environment file with TOML format.")
	err = toml.Encode(b, e, toml.EncodeOption{Indent: &indent})
	if err != nil {
		logger.Logger().Errorf("Error:\n\terr: %v", err)
		return err
	}

	// Write to file
	err = writeFile(path, b.Bytes())
	if err != nil {
		logger.Logger().Errorf("Error:\n\terr: %v", err)
		return err
	}

	logger.Logger().Tracef("Finished")
	return err
}
