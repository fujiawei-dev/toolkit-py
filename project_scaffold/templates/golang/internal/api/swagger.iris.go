{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/iris-contrib/swagger/v12"
	irisSwagger "github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/docs"
)

func init() {
	AddRouteRegistrar(Swagger)
}

func Swagger(router iris.Party) {
	if !conf.Swagger() {
		return
	}

	docs.SwaggerInfo.Title = conf.AppTitle()
	docs.SwaggerInfo.Description = conf.AppDescription()
	docs.SwaggerInfo.Host = conf.ExternalHttpHostPort()
	docs.SwaggerInfo.BasePath = conf.BasePath()
	docs.SwaggerInfo.Version = conf.AppVersion()

	swaggerUI := swagger.WrapHandler(irisSwagger.Handler)

	// Register on /swagger
	router.Get("/swagger", swaggerUI)

	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	router.Get("/swagger/{any:path}", swaggerUI)

	log.Infof("swagger: http://%s/swagger/index.html", conf.ExternalHttpHostPort()+conf.BasePath())
}
