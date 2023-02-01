package systemd

import (
	"bytes"
	"strings"
	"systemd-cd/domain/logger"
	"systemd-cd/domain/toml"
	"systemd-cd/domain/unix"
)

// loadEnvFile implements iSystemdService
func (s Systemd) loadEnvFile(path string) (e map[string]string, isGeneratedBySystemdCd bool, err error) {
	logger.Logger().Debug("START - Load systemd env file")
	logger.Logger().Debugf("< path = %v", path)
	defer func() {
		if err == nil {
			logger.Logger().Debugf("> e = %+v", e)
			logger.Logger().Debugf("> isGeneratedBySystemdCd = %v", isGeneratedBySystemdCd)
			logger.Logger().Debug("END   - Load systemd env file")
		} else {
			logger.Logger().Error("FAILED - Load systemd env file")
			logger.Logger().Error(err)
		}
	}()

	// Read file
	b := &bytes.Buffer{}
	err = unix.ReadFile(path, b)
	if err != nil {
		return
	}

	// Check generator
	if strings.Contains(b.String(), "#! Generated by systemd-cd\n") {
		isGeneratedBySystemdCd = true
	}

	// Decode
	// TODO: toml format で問題ないか検証
	logger.Logger().Warn("Unchecked code: no problem to systemd service environment file with TOML format.")
	err = toml.Decode(b, &e)
	if err != nil {
		return
	}

	return
}
