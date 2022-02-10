{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"net/http"

	"github.com/kataras/iris/v12/core/router"
)

func RemoveTrailingSlash() router.WrapperFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		path := r.URL.Path

		if len(path) > 1 && path[len(path)-1] == '/' && path[len(path)-2] != '/' {
			path = path[:len(path)-1]
			r.RequestURI = path
			r.URL.Path = path
		}

		next(w, r)
	}
}
