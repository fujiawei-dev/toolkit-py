package config

import (
	"{{ main_module }}/internal/event"
)

func (c *config) initLogger() {
	event.SetLevel(c.LogLevelString())
	event.Logger().SetOutput(c.LogWriter())
	event.Logger().SetTimeFormat(c.LogTimeFormat())

	// if c.DetachServer() {
	//	event.Logger().SetFormat("json")
	// }
}

func (c *config) LogLevelString() string {
	if c.settings.Logger.Level != "" {
		return c.settings.Logger.Level
	}

	return "debug"
}
