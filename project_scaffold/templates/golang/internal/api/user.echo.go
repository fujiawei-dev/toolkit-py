{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/labstack/echo/v4"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
}

func PostUser(router *echo.Group) {
	router.POST("/user", func(c echo.Context) error {
		var f form.User

		if err := form.ShouldBind(c, &f); err != nil {
			return ErrorInvalidParameters(c, err)
		}

		if err := entity.CreateWithPassword(f); err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendOK(c)
	})
}

func UserLogin(router *echo.Group) {
	router.POST("/user/login", func(c echo.Context) error {
		var f form.UserLogin

		if err := form.ShouldBind(c, &f); err != nil {
			return ErrorInvalidParameters(c, err)
		}

		m, err := entity.FindUserByUsername(f.Username)
		if err != nil {
			return ErrorExpectedOrUnexpected(c, err)
		}

		if m.InvalidPassword(f.Password) {
			return AbortInvalidPassword(c)

		}

		token, err := conf.JWTGenerate(m)
		if err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendJSON(c, token)
	})
}

func UserLogout(router *echo.Group) {
	router.POST("/user/logout", func(c echo.Context) error {
		user, pass, err := Auth(c, acl.ResourceUsers, acl.ActionUpdate)

		if !pass {
			return err
		}

		log.Infof("user: %s logout", user.Username)

		return SendOK(c)
	}, conf.JWTMiddleware())
}
