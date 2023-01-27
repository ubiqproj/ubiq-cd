package main

import (
	"fmt"
	"os"
	"systemd-cd/domain/model/git"
	"systemd-cd/domain/model/logger"
	"systemd-cd/domain/model/logrus"
	"systemd-cd/domain/model/pipeline"
	"systemd-cd/domain/model/systemd"
	"systemd-cd/infrastructure/externalapi/git_command"
	"systemd-cd/infrastructure/externalapi/systemctl"
	"time"

	logruss "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// flags
var (
	logLevel                  = pflag.String("log.level", "info", "Only log messages with the given severity or above. One of: [panic, fatal, error, warn, info, debug, trace]")
	logReportCaller           = pflag.Bool("log.report-caller", false, "Enable log report caller")
	logTimestamp              = pflag.Bool("log.timestamp", false, "Enable log timestamp.")
	varDir                    = pflag.String("storage.var-dir", "/var/lib/systemd-cd/", "Path to variable files")
	srcDestDir                = pflag.String("storage.src-dir", "/usr/local/systemd-cd/src/", "Path to service source files")
	binaryDestDir             = pflag.String("storage.binary-dir", "/usr/local/systemd-cd/bin/", "Path to service binary files")
	etcDestDir                = pflag.String("storage.etc-dir", "/usr/local/systemd-cd/etc/", "Path to service etc files")
	optDestDir                = pflag.String("storage.opt-dir", "/usr/local/systemd-cd/opt/", "Path to service opt files")
	systemdUnitFileDestDir    = pflag.String("systemd.unit-file-dir", "/usr/local/lib/systemd/system/", "Path to systemd unit files.")
	systemdUnitEnvFileDestDir = pflag.String("systemd.unit-env-file-dir", "/usr/local/systemd-cd/etc/default/", "Path to systemd env files")
	backupDestDir             = pflag.String("storage.backup-dir", "/var/backups/systemd-cd/", "Path to service backup files")
)

func convertLogLevel(str string) (ok bool, lv logger.Level) {
	switch str {
	case "panic":
		return true, logger.PanicLevel
	case "fatal":
		return true, logger.FatalLevel
	case "error":
		return true, logger.ErrorLevel
	case "warn":
		return true, logger.WarnLevel
	case "info":
		return true, logger.InfoLevel
	case "debug":
		return true, logger.DebugLevel
	case "trace":
		return true, logger.TraceLevel
	default:
		return false, logger.InfoLevel
	}
}

func main() {
	// parse flags
	pflag.Parse()

	logger.Init(logrus.New(logrus.Param{
		ReportCaller: logReportCaller,
		Formatter: &logruss.TextFormatter{
			FullTimestamp:   *logTimestamp,
			TimestampFormat: time.RFC3339Nano,
		},
	}))

	// `--log.level`
	ok, lv := convertLogLevel(*logLevel)
	if !ok {
		logger.Logger().Fatalf("`--log.level` must be specified as \"panic\", \"fatal\", \"error\", \"warn\", \"info\", \"debug\" or \"trace\"")
	}
	logger.Logger().SetLevel(lv)

	s, err := systemd.New(systemctl.New(), *systemdUnitFileDestDir)
	if err != nil {
		logger.Logger().Fatalf("Failed:\n\terr: %v", err)
		os.Exit(1)
	}

	g := git.NewService(git_command.New())

	p, err := pipeline.NewService(
		g, s,
		pipeline.Directories{
			Var:                *varDir,
			Src:                *srcDestDir,
			Binary:             *binaryDestDir,
			Etc:                *etcDestDir,
			Opt:                *optDestDir,
			SystemdUnitFile:    *systemdUnitFileDestDir,
			SystemdUnitEnvFile: *systemdUnitEnvFileDestDir,
			Backup:             *backupDestDir,
		},
	)
	if err != nil {
		logger.Logger().Fatalf("Failed:\n\terr: %v", err)
		os.Exit(1)
	}

	p1, err := p.NewPipeline(pipeline.ServiceManifestLocal{
		GitRemoteUrl:    "https://github.com/tingtt/prometheus_sh_exporter.git",
		GitTargetBranch: "main",
		GitManifestFile: nil,
		Name:            "prometheus_sh_exporter",
		Description:     func() *string { s := "The shell exporter allows probing with shell scripts."; return &s }(),
		Port:            func() *uint16 { p := uint16(9923); return &p }(),
		TestCommand:     nil,
		BuildCommand:    func() *string { s := "/usr/bin/go build"; return &s }(),
		Opt:             &[]string{},
		Etc: &[]pipeline.PathOption{{
			Target: "sh.yml",
			Option: "-config.file",
		}},
		Env:            []pipeline.Env{},
		Binary:         func() *string { s := "prometheus_sh_exporter"; return &s }(),
		ExecuteCommand: func() *string { s := "prometheus_sh_exporter"; return &s }(),
		Args:           nil,
	})
	if err != nil {
		logger.Logger().Fatalf("Failed:\n\terr: %v", err)
		os.Exit(1)
	}

	fmt.Printf("p1.GetStatus(): %v\n", p1.GetStatus())
	fmt.Printf("p1.GetCommitRef(): %v\n", p1.GetCommitRef())
}
