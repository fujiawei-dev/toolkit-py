{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"{{GOLANG_MODULE}}/docs"
)

func init() {
	AddRouteRegistrar(Swagger)
}

func Swagger(router fiber.Router) {
	if !conf.Swagger() {
		return
	}

	docs.SwaggerInfo.Title = conf.AppTitle()
	docs.SwaggerInfo.Description = conf.AppDescription()
	docs.SwaggerInfo.Host = conf.ExternalHttpHostPort()
	docs.SwaggerInfo.BasePath = conf.BasePath()
	docs.SwaggerInfo.Version = conf.AppVersion()

	// Register on /swagger
	router.Get("/swagger", swagger.Handler)

	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	router.Get("/swagger/*", swagger.New(swagger.Config{
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "list",
	}))

	log.Info().Msgf("swagger: http://%s/swagger/index.html", conf.ExternalHttpHostPort()+conf.BasePath())
}
