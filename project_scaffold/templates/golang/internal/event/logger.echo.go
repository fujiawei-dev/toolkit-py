{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
)

const TimeFormat = "2006-01-02 15:04:05"

var log echo.Logger

func init() {
	log = echoLog.New("echo")

	log.SetLevel(echoLog.DEBUG)
}

func Logger() echo.Logger {
	return log
}

func SetLevel(l echoLog.Lvl) {
	log.SetLevel(l)
}

func SetLogger(logger echo.Logger) {
	log = logger
}
