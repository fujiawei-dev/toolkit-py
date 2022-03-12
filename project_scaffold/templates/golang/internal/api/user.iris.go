{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
}

func PostUser(router iris.Party) {
	router.Post("/user", func(c iris.Context) {
		var f form.User

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		if err := entity.CreateWithPassword(f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func UserLogin(router iris.Party) {
	router.Post("/user/login", func(c iris.Context) {
		var f form.UserLogin

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		m, err := entity.FindUserByUsername(f.Username)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		if m.InvalidPassword(f.Password) {
			AbortInvalidPassword(c)
			return
		}

		token, err := conf.JWTGenerate(c, m)
		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendJSON(c, token)
	})
}

func UserLogout(router iris.Party) {
	router.Post("/user/logout", conf.JWTMiddleware(), func(c iris.Context) {
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate)

		if !pass {
			return
		}

		log.Infof("user: %s logout", user.Username)

		if err := c.Logout(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}
