{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"crypto/rsa"
	"strings"
	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	"{{GOLANG_MODULE}}/internal/entity"
)

type JWTProvider struct{
	once         sync.Once
	signatureAlg jwt.Alg
	signer       *jwt.Signer
	verifier     *jwt.Verifier
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
}

func (c *config) JWTMiddleware() iris.Handler {
	if !c.JWTEnable() {
		return func(ctx *context.Context) {
			ctx.Next()
		}
	}

	return c.jwtVerifier().Verify(func() interface{} {
		return &entity.User{}
	})
}

func (c *config) JWTGenerate(ctx iris.Context, user entity.User) (interface{}, error) {
	if !c.JWTEnable() {
		return nil, nil
	}

	switch c.JWTMode() {
	case JWTDefault:
		buf, err := c.jwtSigner().Sign(user, jwt.Claims{Issuer: conf.JWTIssuer()})
		value := string(buf)
		ctx.ResponseWriter().Header().Set("Authorization", value)
		return value, err
	case JWTRefresh:
		return c.jwtSigner().NewTokenPair(
			user,
			jwt.Claims{Issuer: conf.JWTIssuer(), Subject: user.Username},
			c.JWTExpire()*8,
			jwt.Claims{Issuer: conf.JWTIssuer()},
		)
	}

	return nil, nil
}

func (c *config) JWTParse(ctx iris.Context) entity.User {
	if !c.JWTEnable() {
		return entity.Admin
	}

	v := jwt.Get(ctx)

	if v == nil {
		return entity.User{}
	}

	user, ok := v.(*entity.User)

	if ok {
		return *user
	}

	return entity.User{}
}

func (c *config) JWTRefresh(ctx iris.Context, user entity.User) (bool, interface{}, error) {
	q := ctx.URLParamTrim("q")

	if len(q) == 0 {
		return false, nil, nil
	}

	_, err := c.jwtVerifier().VerifyToken([]byte(q), jwt.Expected{Issuer: conf.JWTIssuer(), Subject: user.Username})

	if err != nil {
		return false, nil, err
	}

	val, err := c.JWTGenerate(ctx, user)

	return true, val, err
}

func (c *config) initJWT() {
	if !c.JWTEnable() {
		return
	}

	c.settings.JWT.once.Do(func() {
		switch c.JWTMode() {
		case JWTDefault:
			c.settings.JWT.signatureAlg = jwt.HS256
			c.settings.JWT.signer = jwt.NewSigner(c.settings.JWT.signatureAlg, c.JWTKey(), c.JWTExpire())
			c.settings.JWT.verifier = jwt.NewVerifier(c.settings.JWT.signatureAlg, conf.JWTKey(), jwt.Expected{Issuer: conf.JWTIssuer()}).WithDefaultBlocklist()
		case JWTRefresh:
			c.settings.JWT.signatureAlg = jwt.RS256
			c.settings.JWT.privateKey, c.settings.JWT.publicKey = jwt.MustLoadRSA(c.settings.JWT.PrivateKey, c.settings.JWT.PublicKey)
			c.settings.JWT.signer = jwt.NewSigner(c.settings.JWT.signatureAlg, c.settings.JWT.privateKey, c.JWTExpire())
			c.settings.JWT.verifier = jwt.NewVerifier(c.settings.JWT.signatureAlg, c.settings.JWT.publicKey, jwt.Expected{Issuer: conf.JWTIssuer()}).WithDefaultBlocklist()
		}
	})

	// extract token only from Authorization: Bearer $token
	c.settings.JWT.verifier.Extractors = []jwt.TokenExtractor{c.jwtFromHeader}
}

func (c *config) jwtSigner() *jwt.Signer {
	return c.settings.JWT.signer
}

func (c *config) jwtVerifier() *jwt.Verifier {
	return c.settings.JWT.verifier
}

// fromHeader is a token extractor.
// It reads the token from the Authorization request header of form:
// Authorization: "Bearer {token}".
func (c *config) jwtFromHeader(ctx iris.Context) string {
	value := ctx.GetHeader(c.JWTField())
	l := len(c.JWTScheme())
	if len(value) > l+1 && strings.EqualFold(value[:l], c.JWTScheme()) {
		return value[l+1:]
	}
	return ""
}
