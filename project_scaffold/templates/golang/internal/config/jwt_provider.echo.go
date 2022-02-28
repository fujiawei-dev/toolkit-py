{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"{{GOLANG_MODULE}}/internal/entity"
)

type JWTProvider struct {
	once          sync.Once
	jwtMiddleware echo.MiddlewareFunc
}

type jwtUserClaims struct {
	*jwt.StandardClaims
	entity.User
}

func (c *config) JWTMiddleware() echo.MiddlewareFunc {
	if !c.JWTEnable() {
		return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return handlerFunc(ctx)
			}
		}
	}

	c.settings.JWT.once.Do(func() {
		c.settings.JWT.jwtMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:        &jwtUserClaims{},
			ContextKey:    c.JWTContextKey(),
			TokenLookup:   "header:" + c.JWTField(),
			AuthScheme:    c.JWTScheme(),
			SigningMethod: middleware.AlgorithmHS256,
			SigningKey:    c.JWTKey(),
		})
	})

	return c.settings.JWT.jwtMiddleware
}

func (c *config) JWTGenerate(user entity.User) (string, error) {
	claims := jwtUserClaims{
		User: user,
		StandardClaims: &jwt.StandardClaims{
			Issuer:    c.JWTIssuer(),
			ExpiresAt: time.Now().Add(c.JWTExpire()).Unix(),
		},
	}

	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return newClaims.SignedString(c.JWTKey())
}

func (c *config) JWTParse(ctx echo.Context) entity.User {
	user := ctx.Get(c.JWTContextKey()).(*jwt.Token)
	return user.Claims.(*jwtUserClaims).User
}

func (c *config) initJWT() {
	if !c.JWTEnable() {
		return
	}
}
