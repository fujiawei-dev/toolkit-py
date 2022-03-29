{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{WEB_FRAMEWORK_IMPORT}}"
)

func init() {
	AddRouteRegistrar(Debug)
}

func Debug(router {{ROUTER_GROUP}}) {
	router.{{POST_STRING}}("/debug/1", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		SendJSON(c, JsonObject{
			"code": "100",
		})
		return{{NIL_STRING}}
	})

	router.{{POST_STRING}}("/debug/2", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		SendJSON(c, JsonObject{
			"code": "100",
		})
		return{{NIL_STRING}}
	})
}
