package api

import (
	"{{ web_framework_import }}"

	"{{ main_module }}/internal/acl"
	"{{ main_module }}/internal/entity"
)

func Auth(c {{ web_framework_context }}, resource acl.Resource, action acl.Action) (user entity.User, allow bool) {
	defer func() {
		operationLog := entity.NewOperationLog(user.ID, resource, action, allow)
		if err := operationLog.Create(); err != nil {
			log.Printf("create operation log, %v", err)
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

{% if web_framework==".iris" %}
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
{%- endif %}
