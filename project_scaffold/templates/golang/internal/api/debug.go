{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"{{WEB_FRAMEWORK_IMPORT}}"
)

func init() {
	AddRouteRegistrar(Debug)
}

// Debug
// @Summary      请求测试
// @Description  接受任何合法请求类型，返回服务端收到的全部请求参数
// @Tags         功能测试
// @Param        any  path  string  false  "路径参数"
// @Accept       application/x-www-form-urlencoded
// @Produce      plain
// @Success      200,default  {string}  string  "POST / HTTP/1.1\r\nHost: www.example.org\r\nAccept-Encoding: gzip\r\nContent-Length: 75\r\nUser-Agent: Go-http-client/1.1\r\n\r\nGo is a general-purpose language designed with systems programming in mind."
// @Router       /debug/{any} [get]
func Debug(router {{ROUTER_GROUP}}) {
	router.Any("/debug /debug/{any:path}", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		dump, err := httputil.DumpRequest(c.Request(), true)

		if err != nil {
			http.Error(c.ResponseWriter(), fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		c.Write(dump)

		return{{NIL_STRING}}
	})
}
