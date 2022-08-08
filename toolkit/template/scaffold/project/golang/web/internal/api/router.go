package api

import "{{ web_framework_import }}"

type RouteRegistrar func(router {{ web_framework_router_group }})

var RouteRegistrars []RouteRegistrar

func AddRouteRegistrar(rr RouteRegistrar) {
	RouteRegistrars = append(RouteRegistrars, rr)
}

func RegisterRoutes(app {{ web_framework_engine }}) {
	router := app.{{ web_framework_engine_group }}(conf.BasePath())

	for _, rr := range RouteRegistrars {
		rr(router)
	}
}
