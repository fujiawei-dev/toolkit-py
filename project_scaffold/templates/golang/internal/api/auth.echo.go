{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/echo/v4"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Auth(c echo.Context, resource acl.Resource, action acl.Action) (entity.User, bool, error) {
	user := conf.JWTParse(c)

	if user.Invalid() {
		return user, false, AbortUnauthorized(c)
	}

	if acl.Permissions.Deny(resource, user.Role, action) {
		return user, false, AbortPermissionDenied(c)
	}

	return user, true, nil
}
