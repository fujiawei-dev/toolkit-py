{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/echo/v4"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(GetAppDescription)
}

func GetAppDescription(router *echo.Group) {
	router.GET("/app/description", func(c echo.Context) error {
		return SendJSON(c, query.GetAppDescription())
	})
}
