{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func Logger(writer io.Writer, timeFormat string) iris.Handler {
	l := accesslog.New(writer)
	l.TimeFormat = timeFormat
	return l.Handler
}
