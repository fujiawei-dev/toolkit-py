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

        {% if WEB_FRAMEWORK==".iris" -%}
		code := c.URLParam("code")
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		code := c.Query("code")
        {% endif -%}
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
