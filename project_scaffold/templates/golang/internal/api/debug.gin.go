{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	AddRouteRegistrar(Debug)
}

func Debug(router *gin.RouterGroup) {
	router.POST("/debug/1", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "100",
			},
		)
	})

	router.POST("/debug/2", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": "100",
			},
		)
	})
}
