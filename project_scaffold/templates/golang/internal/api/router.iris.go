{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "github.com/kataras/iris/v12"

type routerApi func(router iris.Party)

var routerApis []routerApi

func extendRouterApis(e routerApi) {
	routerApis = append(routerApis, e)
}

func InitRouter(router iris.Party) {
	for _, e := range routerApis {
		e(router)
	}
}
