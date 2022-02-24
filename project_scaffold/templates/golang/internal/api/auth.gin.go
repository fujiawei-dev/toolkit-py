{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Auth(c *gin.Context, resource acl.Resource, action acl.Action) (entity.User, bool) {
	user := conf.JWTParse(c)

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
