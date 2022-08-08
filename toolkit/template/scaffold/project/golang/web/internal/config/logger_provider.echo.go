package config

import (
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"

	"{{ main_module }}/internal/event"
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
