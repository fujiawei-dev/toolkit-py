{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"{{WEB_FRAMEWORK_IMPORT}}"

	"{{GOLANG_MODULE}}/internal/api"
)

func singleJoiningSlash(left, right string) string {
	leftSlash := strings.HasSuffix(left, "/")
	rightSlash := strings.HasPrefix(right, "/")

	switch {
	case leftSlash && rightSlash:
		return left + right[1:]
	case !leftSlash && !rightSlash:
		return left + "/" + right
	}

	return left + right
}

func joinURLPath(left, right *url.URL) (path, rawpath string) {
	if left.RawPath == "" && right.RawPath == "" {
		return singleJoiningSlash(left.Path, right.Path), ""
	}

	leftPath := left.EscapedPath()
	rightPath := right.EscapedPath()

	leftSlash := strings.HasSuffix(leftPath, "/")
	rightSlash := strings.HasPrefix(rightPath, "/")

	switch {
	case leftSlash && rightSlash:
		return left.Path + right.Path[1:], leftPath + rightPath[1:]
	case !leftSlash && !rightSlash:
		return left.Path + "/" + right.Path, leftPath + "/" + rightPath
	}

	return left.Path + right.Path, leftPath + rightPath
}

func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host

		// req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
		req.URL.Path, req.URL.RawPath = target.Path, target.RawPath

		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}

		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "golang/1.18.1")
		}
	}

	return &httputil.ReverseProxy{Director: director}
}

func NewReverseProxyFunc(targetUrl string, clientIP string) (http.HandlerFunc, error) {
	if !strings.HasPrefix(targetUrl, "http") {
		targetUrl = "http://" + targetUrl
	}

	target, err := url.Parse(targetUrl)
	if err != nil {
		return nil, err
	}

	reverseProxy := NewReverseProxy(target)

	reverseProxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Set("Client-IP", clientIP)
		r.Header.Set("X-Proxy", "Reverse Proxy")
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	}, nil
}

// NewReverseProxyIrisExample
// @Summary      逆向代理测试
// @Description  逆向代理测试
// @Tags         功能测试
// @Accept       application/x-www-form-urlencoded
// @Produce      plain
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /proxy [get]
func NewReverseProxyIrisExample(targetUrl string, router iris.Party, path string) {
	router.Any(path, func(c iris.Context) {
		proxy, err := NewReverseProxyFunc(targetUrl, c.RemoteAddr())
		if err != nil {
			api.ErrorUnexpected(c, err)
			return
		}

		proxy(c.ResponseWriter(), c.Request())
	})
}
