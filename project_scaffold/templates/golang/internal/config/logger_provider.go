{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"{{GOLANG_MODULE}}/internal/event"
)

var (
	writer     io.Writer = os.Stdout
	writerOnce sync.Once
)

func (c *config) initLogger() {
	writerOnce.Do(func() {
		if c.DetachServer() {
			event.SetLogger(zerolog.New(c.LogWriter()).
				With().
				CallerWithSkipFrameCount(3).
				Logger(),
			)

			event.SetLevel(c.LogLevel())
		}
	})
}

func (c *config) LogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(c.settings.Logger.Level)

	if err != nil {
		return zerolog.InfoLevel
	}

	return level
}

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
