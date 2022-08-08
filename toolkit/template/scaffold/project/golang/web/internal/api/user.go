package api

import (
	"errors"

	"{{ web_framework_import }}"
    {%- if web_framework not in [".iris"] %}
	"github.com/spf13/cast"
    {%- endif %}

	"{{ main_module }}/internal/acl"
	"{{ main_module }}/internal/entity"
	"{{ main_module }}/internal/form"
	"{{ main_module }}/internal/query"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
	AddRouteRegistrar(PutUser)
	AddRouteRegistrar(PutUserPassword)
	AddRouteRegistrar(GetUsers)
}

func PostUser(router {{ web_framework_router_group }}) {
	router.{{ web_framework_post }}("/user", func(c {{ web_framework_context }}) {{ web_framework_error }}{
		var f form.User

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		if err := entity.CreateWithPassword(f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	})
}

func PutUser(router {{ web_framework_router_group }}) {
	router.{{ web_framework_put }}("/user/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate); !pass {
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif web_framework==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif web_framework==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif web_framework==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{ web_framework_nil }}
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		// Handle null values, malicious injection, etc.
		var f form.UserUpdate

		if err = m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		if err = m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err = m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}

func PutUserPassword(router {{ web_framework_router_group }}) {
	router.{{ web_framework_put }}("/user/{{ web_framework_query_id }}/password", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourcePasswords, acl.ActionUpdate); !pass {
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif web_framework==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif web_framework==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif web_framework==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{ web_framework_nil }}
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		// Handle null values, malicious injection, etc.
		var f form.UserChangePassword

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		if m.InvalidPassword(f.OldPassword) {
			AbortInvalidPassword(c)
			return{{ web_framework_nil }}
		}

		if err = m.SetPassword(f.NewPassword); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}

func GetUsers(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/users", func(c {{ web_framework_context }}) {{ web_framework_error }}{
		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		username := c.URLParam("username")
        {% elif web_framework in [".gin", ".fiber"] -%}
		username := c.Query("username")
        {% elif web_framework==".echo" -%}
		username := c.QueryParam("username")
        {% endif -%}
		f.LikeQ = username

		list, totalRow, err := query.Users(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{ web_framework_nil }}
	})
}

func UserLogin(router {{ web_framework_router_group }}) {
	router.{{ web_framework_post }}("/user/login", func(c {{ web_framework_context }}) {{ web_framework_error }}{
		var f form.UserLogin

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		m, err := entity.FindUserByUsername(f.Username)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		defer func() {
			operationLog := entity.NewOperationLog(m.ID, acl.ResourceUsers, acl.ActionLogin, err == nil)
			if err = operationLog.Create(); err != nil {
				log.Printf("create operation log, %v", err)
			}
		}()

		if m.InvalidPassword(f.Password) {
			AbortInvalidPassword(c)
			return{{ web_framework_nil }}
		}

		token, err := conf.JWTGenerate(m)
		if err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		c.Header("Authorization", token.(string))
		c.Header("Access-Control-Expose-Headers", "Authorization")
        {% elif web_framework==".fiber" -%}
		c.Response().Header.Set("Authorization", token)
		c.Response().Header.Set("Access-Control-Expose-Headers", "Authorization")
        {% elif web_framework==".echo" -%}
		c.Response().Header().Set("Authorization", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "Authorization")
        {% elif web_framework==".gin" -%}
		c.Header("Authorization", token)
		c.Header("Access-Control-Expose-Headers", "Authorization")
        {% endif -%}

		var r query.UserResult
		m.CopyTo(&r)

		SendJSON(c, r)
		return{{ web_framework_nil }}
	})
}

func UserLogout(router {{ web_framework_router_group }}) {
	router.{{ web_framework_post }}("/user/logout", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionLogout)

		if !pass {
			return{{ web_framework_nil }}
		}

		log.Printf("user: %s logout", user.Username)

		SendOK(c)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}
