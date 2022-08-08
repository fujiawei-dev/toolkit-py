package config

import (
	"github.com/rs/zerolog"

	"{{ main_module }}/internal/event"
)


func (c *config) initLogger() {
	if c.DetachServer() {
		event.SetLogger(zerolog.New(c.LogWriter()).
			With().
			CallerWithSkipFrameCount(2).
			Logger(),
		)

		event.SetLevel(c.LogLevel())
	}
}

func (c *config) LogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(c.settings.Logger.Level)

	if err != nil {
		return zerolog.InfoLevel
	}

	return level
}
