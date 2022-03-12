{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(GetAppDescription)
}

func GetAppDescription(router iris.Party) {
	router.Get("/app/description", func(c iris.Context) {
		SendJSON(c, query.GetAppDescription())
	})
}
