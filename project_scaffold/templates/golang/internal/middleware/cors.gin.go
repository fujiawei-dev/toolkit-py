{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	// I just need to know whether it is safe or not, as for whether it is actually safe or not, it has nothing to do with me.
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Cache-Control", "Authorization",
			"Content-Language", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
