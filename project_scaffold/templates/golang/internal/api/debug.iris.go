{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/internal/query"
)

func RegisterDebug(router iris.Party) {
	if !conf.Debug() {
		return
	}

	router.Get("/debug", func(c iris.Context) {
		SendJSON(c, query.GetAppVersion())
	})
}
