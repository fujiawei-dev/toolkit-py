package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"{{ web_framework_import }}"
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
func Debug(router {{ web_framework_router_group }}) {
    {% if web_framework==".iris" -%}
	router.Any("/debug /debug/{any:path}", func(c {{ web_framework_context }}) {{ web_framework_error }}{
		dump, err := httputil.DumpRequest(c.Request(), true)
		if err != nil {
			http.Error(c.ResponseWriter(), fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		c.Write(dump)
    id := c.Params().GetUintDefault("id", 0)
    {% elif web_framework==".fiber" -%}
    id := cast.ToUint(c.Params("id"))
    {% elif web_framework==".echo" -%}
    id := cast.ToUint(c.Param("id"))
    {% elif web_framework==".gin" -%}
	router.Any("/debug /debug/{any:path}", func(c *gin.Context) {
		dump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			http.Error(c.Writer, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		c.Writer.Write(dump)
    {% endif -%}

		return{{ web_framework_nil }}
	})
}
