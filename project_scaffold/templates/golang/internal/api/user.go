{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"

	"{{WEB_FRAMEWORK_IMPORT}}"
    {%- if WEB_FRAMEWORK not in [".iris"] %}
	"github.com/spf13/cast"
    {%- endif %}

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
	AddRouteRegistrar(PutUser)
	AddRouteRegistrar(PutUserPassword)
	AddRouteRegistrar(GetUsers)
}

func PostUser(router {{ROUTER_GROUP}}) {
	router.{{POST_STRING}}("/user", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		var f form.User

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		if err := entity.CreateWithPassword(f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	})
}

func PutUser(router {{ROUTER_GROUP}}) {
	router.{{PUT_STRING}}("/user/{{QUERY_ID}}", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate); !pass {
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{NIL_STRING}}
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		// Handle null values, malicious injection, etc.
		var f form.UserUpdate

		if err = m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		if err = m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err = m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

func PutUserPassword(router {{ROUTER_GROUP}}) {
	router.{{PUT_STRING}}("/user/{{QUERY_ID}}/password", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourcePasswords, acl.ActionUpdate); !pass {
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{NIL_STRING}}
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		// Handle null values, malicious injection, etc.
		var f form.UserChangePassword

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		if m.InvalidPassword(f.OldPassword) {
			AbortInvalidPassword(c)
			return{{NIL_STRING}}
		}

		if err = m.SetPassword(f.NewPassword); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

func GetUsers(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/users", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		username := c.URLParam("username")
        {% elif WEB_FRAMEWORK in [".gin", ".fiber"] -%}
		username := c.Query("username")
        {% elif WEB_FRAMEWORK==".echo" -%}
		username := c.QueryParam("username")
        {% endif -%}
		f.LikeQ = username

		list, totalRow, err := query.Users(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{NIL_STRING}}
	})
}

func UserLogin(router {{ROUTER_GROUP}}) {
	router.{{POST_STRING}}("/user/login", func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		var f form.UserLogin

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		m, err := entity.FindUserByUsername(f.Username)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		defer func() {
			operationLog := entity.NewOperationLog(m.ID, acl.ResourceUsers, acl.ActionLogin, err == nil)
			if err = operationLog.Create(); err != nil {
				log.Printf("create operation log, %v", err)
			}
		}()

		if m.InvalidPassword(f.Password) {
			AbortInvalidPassword(c)
			return{{NIL_STRING}}
		}

		token, err := conf.JWTGenerate(m)
		if err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		c.Header("Authorization", token.(string))
		c.Header("Access-Control-Expose-Headers", "Authorization")
        {% elif WEB_FRAMEWORK==".fiber" -%}
		c.Response().Header.Set("Authorization", token)
		c.Response().Header.Set("Access-Control-Expose-Headers", "Authorization")
        {% elif WEB_FRAMEWORK==".echo" -%}
		c.Response().Header().Set("Authorization", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "Authorization")
        {% elif WEB_FRAMEWORK==".gin" -%}
		c.Header("Authorization", token)
		c.Header("Access-Control-Expose-Headers", "Authorization")
        {% endif -%}

		var r query.UserResult
		m.CopyTo(&r)

		SendJSON(c, r)
		return{{NIL_STRING}}
	})
}

func UserLogout(router {{ROUTER_GROUP}}) {
	router.{{POST_STRING}}("/user/logout", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionLogout)

		if !pass {
			return{{NIL_STRING}}
		}

		log.Printf("user: %s logout", user.Username)

		SendOK(c)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}
