{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

func init() {
	AddRouteRegistrar(UserLogin)
	AddRouteRegistrar(UserLogout)

	AddRouteRegistrar(PostUser)
}

func PostUser(router *gin.RouterGroup) {
	router.POST("/user", func(c *gin.Context) {
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

func UserLogin(router *gin.RouterGroup) {
	router.POST("/user/login", func(c *gin.Context) {
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

		token, err := conf.JWTGenerate(m)
		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendJSON(c, token)
	})
}

func UserLogout(router *gin.RouterGroup) {
	router.POST("/user/logout", conf.JWTMiddleware(), func(c *gin.Context) {
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate)

		if !pass {
			return
		}

		log.Info().Msgf("user: %s logout", user.Username)

		SendOK(c)
	})
}
