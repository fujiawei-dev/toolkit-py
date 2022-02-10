{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/golog"
)

// https://github.com/kataras/golog.git
var Log *golog.Logger

const TimeFormat = "2006-01-02 15:04:05"

func init() {
	Log = golog.New()

	// https://github.com/kataras/golog/issues/3
	Log.SetLevel("debug")
	Log.SetTimeFormat(TimeFormat)
	Log.SetStacktraceLimit(2)
}
