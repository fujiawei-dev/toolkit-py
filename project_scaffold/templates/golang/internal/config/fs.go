{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/pkg/fs"
)

func (c *config) CreateDirectories() (err error) {
	if !fs.Exists(c.StoragePath()) {
		if err = fs.MkdirAll(c.StoragePath()); err != nil {
			return err
		}
	}

	if !fs.Exists(c.BackupPath()) {
		if err = fs.MkdirAll(c.BackupPath()); err != nil {
			return err
		}
	}

	if !fs.Exists(c.LogPath()) {
		if err = fs.MkdirAll(c.LogPath()); err != nil {
			return err
		}
	}

	if !fs.Exists(c.ConfigsPath()) {
		if err = fs.MkdirAll(c.ConfigsPath()); err != nil {
			return err
		}
	}

	return nil
}

func (c *config) SettingsFile() string {
	if c.flags.ConfigFile != "" {
		return c.flags.ConfigFile
	}

	return "settings.yaml"
}

func (c *config) ConfigsPath() string {
	return fs.Join(c.StoragePath(), "configs")
}

func (c *config) StoragePath() string {
	if c.settings.Static.StoragePath != "" {
		return fs.MustAbs(c.settings.Static.StoragePath)
	}

	return fs.MustAbs("storage")
}

// BackupPath returns the backup storage path.
func (c *config) BackupPath() string {
	if fs.PathWritable(c.settings.Static.BackupPath) {
		return fs.MustAbs(c.settings.Static.BackupPath)
	}

	return fs.Join(c.StoragePath(), "backup")
}

// PidFile returns the filename for storing the server process id (pid).
func (c *config) PidFile() string {
	return fs.Join(c.StoragePath(), c.AppName()+".pid")
}
