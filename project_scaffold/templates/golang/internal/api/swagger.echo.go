{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"{{GOLANG_MODULE}}/docs"
)

func init() {
	AddRouteRegistrar(Swagger)
}

func Swagger(router *echo.Group) {
	if !conf.Swagger() {
		return
	}

	docs.SwaggerInfo.Title = conf.AppTitle()
	docs.SwaggerInfo.Description = conf.AppDescription()
	docs.SwaggerInfo.Host = conf.ExternalHttpHostPort()
	docs.SwaggerInfo.BasePath = conf.BasePath()
	docs.SwaggerInfo.Version = conf.AppVersion()

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Infof("swagger: http://%s/swagger/index.html", conf.ExternalHttpHostPort()+conf.BasePath())
}
