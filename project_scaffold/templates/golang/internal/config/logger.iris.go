{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/pkg/fs"
)

var logWriter io.Writer = os.Stdout
var logWriterOnce sync.Once

func (c *config) LogWriter() io.Writer {
	logWriterOnce.Do(func() {
		if c.DetachServer() {
			logWriter = &lumberjack.Logger{
				Filename:   conf.LogFile(),
				MaxSize:    conf.LogMaxSize(),
				MaxAge:     conf.LogMaxAge(),
				MaxBackups: conf.LogMaxBackups(),
				LocalTime:  true,
			}
		}
	})

	return logWriter
}

func (c *config) initLogger() {
	event.Log.SetLevel(c.LogLevelString())
	event.Log.SetOutput(c.LogWriter())
	event.Log.SetTimeFormat(c.LogTimeFormat())

	// 这个日志库输出格式很无语，关键信息缺失，必须自己实现，另外 JSON 输出多行而不是一行

	//if c.DetachServer() {
	//	event.Log.SetFormat("json")
	//	event.Log.NewLine = false
	//}
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

func (c *config) LogLevelString() string {
	if c.settings.Logger.Level != "" {
		return c.settings.Logger.Level
	}

	return "info"
}

func (c *config) LogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(c.settings.Logger.Level)

	if err != nil {
		return zerolog.InfoLevel
	}

	return level
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
