{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
}

func PostUser(router fiber.Router) {
	router.Post("/user", func(c *fiber.Ctx) error {
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

func UserLogin(router fiber.Router) {
	router.Post("/user/login", func(c *fiber.Ctx) error {
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

		token, err := conf.JWTGenerate(c, m)
		if err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendJSON(c, token)
	})
}

func UserLogout(router fiber.Router) {
	router.Post("/user/logout", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		user, pass, err := Auth(c, acl.ResourceUsers, acl.ActionUpdate)

		if !pass {
			return err
		}

		log.Info().Msgf("user: %s logout", user.Username)

		return SendOK(c)
	})
}
