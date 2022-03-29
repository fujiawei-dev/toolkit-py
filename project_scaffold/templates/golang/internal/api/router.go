{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "{{WEB_FRAMEWORK_IMPORT}}"

type RouteRegistrar func(router {{ROUTER_GROUP}})

var RouteRegistrars []RouteRegistrar

func AddRouteRegistrar(rr RouteRegistrar) {
	RouteRegistrars = append(RouteRegistrars, rr)
}

func RegisterRoutes(app {{WEB_ENGINE}}) {
	router := app.{{WEB_ENGINE_GROUP}}(conf.BasePath())

	for _, rr := range RouteRegistrars {
		rr(router)
	}
}
