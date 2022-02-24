{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/event"
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
