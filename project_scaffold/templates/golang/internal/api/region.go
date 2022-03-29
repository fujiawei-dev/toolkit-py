{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{WEB_FRAMEWORK_IMPORT}}"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(RegionBy)
}

func RegionBy(router {{ROUTER_GROUP}}) {
	router.Any("/region/by", func(c {{WEB_CONTEXT}}) {
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
		return{{NIL_STRING}}
	})
}
