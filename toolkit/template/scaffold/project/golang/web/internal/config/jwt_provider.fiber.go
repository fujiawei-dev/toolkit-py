package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
	jwtProvider "github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"

	"{{ main_module }}/internal/entity"
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

func (c *config) JWTGenerate(user entity.User) (string, error) {
	if !c.JWTEnable() {
		return "", nil
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

	return "", nil
}

func (c *config) JWTParse(ctx *fiber.Ctx) (user entity.User) {
	if !c.JWTEnable() {
		return entity.Admin
	}

	claims := ctx.Locals(c.JWTContextKey()).(*jwtProvider.Token).Claims.(jwtProvider.MapClaims)

	if v := claims["user"]; v != nil {
		_ = mapstructure.Decode(v, &user)
	}

	return
}

func (c *config) initJWT() {
	if !c.JWTEnable() {
		return
	}
}
