{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/query"
)

func GetAppDescription(router *gin.RouterGroup) {
	router.GET("/app/description", func(c *gin.Context) {
		SendJSON(c, query.GetAppDescription())
	})
}
