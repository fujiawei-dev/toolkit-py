{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
	jwtProvider "github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"

	"{{GOLANG_MODULE}}/internal/entity"
)

type JWTProvider struct{

}

func (c *config) JWTMiddleware() fiber.Handler {
	// https://github.com/gofiber/jwt
	if !c.JWTEnable() {
		return func(ctx *fiber.Ctx) error {
			return ctx.Next()
		}
	}

	return jwt.New(jwt.Config{
		SigningKey: conf.JWTKey(),
		ContextKey: conf.JWTContextKey(),
		AuthScheme: conf.JWTScheme(),
	})
}

func (c *config) JWTGenerate(ctx *fiber.Ctx, user entity.User) (interface{}, error) {
	if !c.JWTEnable() {
		return nil, nil
	}

	switch c.JWTMode() {
	case JWTDefault:
		claims := jwtProvider.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		}
		j := jwtProvider.NewWithClaims(jwtProvider.SigningMethodHS256, claims)
		return j.SignedString(c.JWTKey())
	}

	return nil, nil
}

func (c *config) JWTParse(ctx *fiber.Ctx) (user entity.User, err error) {
	if !c.JWTEnable() {
		return entity.Admin, nil
	}

	claims := ctx.Locals(c.JWTContextKey()).(*jwtProvider.Token).Claims.(jwtProvider.MapClaims)

	if v := claims["user"]; v != nil {
		err = mapstructure.Decode(v, &user)
	}

	return
}

func (c *config) initJWT() {
	if !c.JWTEnable() {
		return
	}
}
