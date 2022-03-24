{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

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

func PutUser(router *gin.RouterGroup) {
	router.PUT("/user/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceUsers, acl.ActionUpdate); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		// Handle null values, malicious injection, etc.
		var f form.UserUpdate

		if err = m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		if err = m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		if err = m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func PutUserPassword(router *gin.RouterGroup) {
	router.PUT("/user/:id/password", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourcePasswords, acl.ActionUpdate); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		m, err := entity.FindUserByID(id)
		if err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		// Handle null values, malicious injection, etc.
		var f form.UserChangePassword

		if err = form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		if m.InvalidPassword(f.OldPassword) {
			AbortInvalidPassword(c)
			return
		}

		if err = m.SetPassword(f.NewPassword); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func GetUsers(router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		f.LikeQ = c.Query("username")

		list, totalRow, err := query.Users(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
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

		defer func() {
			operationLog := entity.NewOperationLog(m.ID, acl.ResourceUsers, acl.ActionLogin, err == nil)
			if err = operationLog.Create(); err != nil {
				log.Error().Msgf("create operation log, %v", err)
			}
		}()

		if m.InvalidPassword(f.Password) {
			AbortInvalidPassword(c)
			return
		}

		token, err := conf.JWTGenerate(m)
		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		c.Header("Authorization", token)
		c.Header("Access-Control-Expose-Headers", "Authorization")

		SendJSON(c, m)
	})
}

func UserLogout(router *gin.RouterGroup) {
	router.POST("/user/logout", conf.JWTMiddleware(), func(c *gin.Context) {
		user, pass := Auth(c, acl.ResourceUsers, acl.ActionLogout)

		if !pass {
			return
		}

		log.Info().Msgf("user: %s logout", user.Username)

		SendOK(c)
	})
}
