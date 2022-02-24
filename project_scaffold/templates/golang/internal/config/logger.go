{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io"
	"os"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/pkg/fs"
)

var (
	writer     io.Writer = os.Stdout
	writerOnce sync.Once
)

func (c *config) LogWriter() io.Writer {
	writerOnce.Do(func() {
		if c.DetachServer() {
			writer = &lumberjack.Logger{
				Filename:   conf.LogFile(),
				MaxSize:    conf.LogMaxSize(),
				MaxAge:     conf.LogMaxAge(),
				MaxBackups: conf.LogMaxBackups(),
				LocalTime:  true,
			}
		}
	})

	return writer
}

func (c *config) LogPath() string {
	if c.settings.Logger.SavePath != "" {
		return fs.MustAbs(c.settings.Logger.SavePath)
	}

	return fs.Join(c.StoragePath(), "logs")
}

func (c *config) LogFile() string {
	if c.settings.Logger.Filename != "" {
		return fs.Join(c.LogPath(), fs.Base(c.settings.Logger.Filename))
	}

	return fs.Join(c.LogPath(), c.AppName()+".log")
}

func (c *config) LogMaxSize() int {
	if c.settings.Logger.MaxSize > 0 {
		return c.settings.Logger.MaxSize
	}

	return 100
}

func (c *config) LogMaxAge() int {
	if c.settings.Logger.MaxAge > 0 {
		return c.settings.Logger.MaxAge
	}

	return 30
}

func (c *config) LogMaxBackups() int {
	if c.settings.Logger.MaxBackups > 0 {
		return c.settings.Logger.MaxBackups
	}

	return 15
}

func (c *config) LogTimeFormat() string {
	if c.DetachServer() {
		return time.RFC3339Nano
	}

	return event.TimeFormat
}
