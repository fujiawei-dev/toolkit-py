package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"{{ main_module }}/docs"
)

func init() {
	AddRouteRegistrar(Swagger)
}

func Swagger(router *gin.RouterGroup) {
	if !conf.Swagger() {
		return
	}

	docs.SwaggerInfo.Title = conf.AppTitle()
	docs.SwaggerInfo.Description = conf.AppDescription()
	docs.SwaggerInfo.Host = conf.ExternalHttpHostPort()
	docs.SwaggerInfo.BasePath = conf.BasePath()
	docs.SwaggerInfo.Version = conf.AppVersion()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("swagger: http://%s/swagger/index.html", conf.ExternalHttpHostPort()+conf.BasePath())
}
