{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/middleware"
)

type JWTProvider struct {
	once          sync.Once
	jwtMiddleware gin.HandlerFunc
}

type jwtUserClaims struct {
	*jwt.StandardClaims
	entity.User
}

func (c *config) JWTMiddleware() gin.HandlerFunc {
	if !c.JWTEnable() {
		return func(ctx *gin.Context) {
			ctx.Next()
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

func (c *config) JWTParse(ctx *gin.Context) entity.User {
	if user, exists := ctx.Get(c.JWTContextKey()); exists {
		return user.(*jwt.Token).Claims.(*jwtUserClaims).User
	}

	return entity.User{}
}

func (c *config) initJWT() {
	if !c.JWTEnable() {
		return
	}
}
