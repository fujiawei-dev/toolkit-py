{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Auth(c *gin.Context, resource acl.Resource, action acl.Action) (user entity.User, allow bool) {
	defer func() {
		operationLog := entity.NewOperationLog(user.ID, resource, action, allow)
		if err := operationLog.Create(); err != nil {
			log.Error().Msgf("create operation log, %v", err)
		}
	}()

	user = conf.JWTParse(c)

	allow = !user.Invalid()
	if !allow {
		AbortUnauthorized(c)
		return
	}

	allow = !acl.Permissions.Deny(resource, user.Role, action)
	if !allow {
		AbortPermissionDenied(c)
		return
	}

	return
}
