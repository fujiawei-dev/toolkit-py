{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(RegionBy)
}

func RegionBy(router *gin.RouterGroup) {
	router.Any("/region/by", func(c *gin.Context) {
		var result interface{}

		code := c.Query("code")
		if code != "" {
			region, location := conf.GetRegionByCode(code)
			result = query.RegionLocation{
				Region:   region,
				Location: location,
			}
		}

		SendJSON(c, result)
	})
}
