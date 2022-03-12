{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "github.com/gin-gonic/gin"

type RouteRegistrar func(router *gin.RouterGroup)

var RouteRegistrars []RouteRegistrar

func AddRouteRegistrar(rr RouteRegistrar) {
	RouteRegistrars = append(RouteRegistrars, rr)
}

func RegisterRoutes(app *gin.Engine) {
	router := app.Group(conf.BasePath())

	for _, rr := range RouteRegistrars {
		rr(router)
	}
}
