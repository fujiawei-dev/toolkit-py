{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
)

func Auth(c *fiber.Ctx, resource acl.Resource, action acl.Action) (entity.User, bool, error) {
	user, err := conf.JWTParse(c)

	if err != nil {
		return user, false, ErrorExpectedOrUnexpected(c, err)
	}

	if user.Invalid() {
		return user, false, AbortUnauthorized(c)
	}

	if acl.Permissions.Deny(resource, user.Role, action) {
		return user, false, AbortPermissionDenied(c)
	}

	return user, true, nil
}
