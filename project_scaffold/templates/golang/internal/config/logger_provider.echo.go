{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"

	"{{GOLANG_MODULE}}/internal/event"
)

func (c *config) initLogger() {
	event.SetLevel(log.Lvl(c.LogLevel() + 1))
	event.Logger().SetOutput(c.LogWriter())
}

func (c *config) LogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(c.settings.Logger.Level)

	if err != nil {
		return zerolog.InfoLevel
	}

	return level
}
