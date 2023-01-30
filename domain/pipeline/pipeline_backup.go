package pipeline

import (
	"strconv"
	"systemd-cd/domain/logger"
	"systemd-cd/domain/unix"
	"time"
)

func (p pipeline) backupInstalled() error {
	logger.Logger().Trace(logger.Var2Text("Called", []logger.Var{{Value: p}}))

	// Create directory for backup
	// e.g.
	// /var/backups/systemd-cd/<name>/<unix-time>_<commit-id>/
	backupPath := p.service.PathBackupDir + p.ManifestMerged.Name + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + p.GetCommitRef() + "/"
	err := unix.MkdirIfNotExist(backupPath)
	if err != nil {
		logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
		return err
	}

	if p.ManifestMerged.SystemdOptions != nil && len(p.ManifestMerged.SystemdOptions) != 0 {
		err = unix.MkdirIfNotExist(backupPath+"systemd/", backupPath+"env/")
		if err != nil {
			logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
			return err
		}
		for _, s := range p.ManifestMerged.SystemdOptions {
			// Backup systemd unit file
			// e.g.
			// `cp /usr/local/lib/systemd/system/<unit_name>.service /var/backups/systemd-cd/<name>/<unix-time>_<commit-id>/systemd/<unit_name>.service`
			err = unix.Mv(
				unix.ExecuteOption{},
				unix.MvOption{},
				p.service.PathSystemdUnitFileDir+s.Name+".service",
				backupPath+"systemd/",
			)
			if err != nil {
				logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
				return err
			}

			if len(s.Opt) != 0 {
				// Backup opt files
				// e.g.
				// `cp /usr/local/systemd-cd/opt/<unit_name>/* /var/backups/systemd-cd/<name>/<unix-time>_<commit-id>/opt/<unit_name>/`
				err = unix.MkdirIfNotExist(backupPath + "opt/" + s.Name)
				if err != nil {
					logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
					return err
				}
				err = unix.Mv(
					unix.ExecuteOption{},
					unix.MvOption{},
					p.service.PathOptDir+s.Name+"/*",
					backupPath+"opt/"+s.Name,
				)
				if err != nil {
					logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
					return err
				}
			}

			// Backup env file
			// e.g.
			// `cp /usr/local/systemd-cd/etc/default/<unit_name> /var/backups/systemd-cd/<name>/<unix-time>_<commit-id>/env/<unit_name>`
			err = unix.Mv(
				unix.ExecuteOption{},
				unix.MvOption{},
				p.service.PathSystemdUnitEnvFileDir+s.Name,
				backupPath+"env/",
			)
			if err != nil {
				logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
				return err
			}
		}
	}

	if p.ManifestMerged.Binaries != nil && len(*p.ManifestMerged.Binaries) != 0 {
		// Backup binary files
		// e.g.
		// `cp /usr/local/systemd-cd/bin/<name>/* /var/backups/systemd-cd/<name>/<unix-time>_<commit-id>/bin/`
		err = unix.MkdirIfNotExist(backupPath + "bin/")
		if err != nil {
			logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
			return err
		}
		err = unix.Mv(
			unix.ExecuteOption{},
			unix.MvOption{},
			p.service.PathBinDir+p.ManifestMerged.Name+"/*",
			backupPath+"bin/",
		)
		if err != nil {
			logger.Logger().Error(logger.Var2Text("Error", []logger.Var{{Name: "err", Value: err}}))
			return err
		}
	}

	logger.Logger().Trace("Finished")
	return nil
}
