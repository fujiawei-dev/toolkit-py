{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Cache-Control",
			"Content-Language", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
