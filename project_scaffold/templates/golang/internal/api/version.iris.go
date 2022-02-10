{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/internal/query"
)

func GetVersion(router iris.Party) {
	router.Get("/system/version", func(c iris.Context) {
		SendJSON(c, query.GetAppVersion())
	})
}
