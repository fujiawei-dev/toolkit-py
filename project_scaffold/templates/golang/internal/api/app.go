{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{WEB_FRAMEWORK_IMPORT}}"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(GetAppDescription)
}

func GetAppDescription(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/app/description", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		SendJSON(c, query.GetAppDescription())
		return{{NIL_STRING}}
	})
}
