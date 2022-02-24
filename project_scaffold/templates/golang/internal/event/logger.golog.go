{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/golog"
)

const TimeFormat = "2006-01-02 15:04:05"

// https://github.com/kataras/golog.git
var log *golog.Logger

func init() {
	log = golog.New()

	// https://github.com/kataras/golog/issues/3
	log.SetLevel("debug")
	log.SetTimeFormat(TimeFormat)
	log.SetStacktraceLimit(2)
}

func Logger() *golog.Logger {
	return log
}

func SetLevel(l string) {
	log.SetLevel(l)
}

func SetLogger(logger *golog.Logger) {
	log = logger
}
