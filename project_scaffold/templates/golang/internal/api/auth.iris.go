{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Auth(c iris.Context, resource acl.Resource, action acl.Action) (user entity.User, pass bool) {
	user = conf.JWTParse(c)

	if user.Invalid() {
		AbortUnauthorized(c)
		return user, false
	}

	if acl.Permissions.Deny(resource, user.Role, action) {
		AbortPermissionDenied(c)
		return user, false
	}

	return user, true
}

// UserLoginRefresh Refresh Token
func UserLoginRefresh(router iris.Party) {
	router.Get("/user/login/refresh", conf.JWTMiddleware(), func(c iris.Context) {
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate)

		if !pass {
			return
		}

		valid, token, err := conf.JWTRefresh(c, user)

		if err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		if !valid {
			AbortUnauthorized(c)
			return
		}

		SendJSON(c, token)
	})
}
