package api

import (
	"{{ web_framework_import }}"

	"{{ main_module }}/internal/query"
)

func init() {
	AddRouteRegistrar(GetAppDescription)
}

func GetAppDescription(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/app/description", func(c {{ web_framework_context }}) {{ web_framework_error }}{
		SendJSON(c, query.GetAppDescription())
		return{{ web_framework_nil }}
	})
}
