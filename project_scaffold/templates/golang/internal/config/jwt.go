{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"{{GOLANG_MODULE}}/pkg/fs"
)

const (
	JWTDefault = "default"
	JWTRefresh = "refresh"
)

func (c *config) JWTEnable() bool {
	return c.settings.JWT.Enable
}

func (c *config) JWTKey() []byte {
	if c.settings.JWT.Key != "" {
		return []byte(c.settings.JWT.Key)
	}

	return []byte("secret")
}

func (c *config) JWTIssuer() string {
	if c.settings.JWT.Issuer != "" {
		return c.settings.JWT.Issuer
	}

	return "issuer"
}

func (c *config) JWTScheme() string {
	if c.settings.JWT.Scheme != "" {
		return c.settings.JWT.Scheme
	}

	return "Bearer"
}

func (c *config) JWTField() string {
	if c.settings.JWT.Field != "" {
		return c.settings.JWT.Field
	}

	return "Authorization"
}

func (c *config) JWTExpire() time.Duration {
	if c.settings.JWT.Expire > time.Hour {
		return c.settings.JWT.Expire
	}

	return time.Hour
}

func (c *config) JWTContextKey() string {
	return "user"
}

func (c *config) JWTMode() string {
	if !fs.IsFile(c.settings.JWT.PrivateKey) ||
		!fs.IsFile(c.settings.JWT.PublicKey) {
		return JWTDefault
	}

	switch c.settings.JWT.Mode {
	case JWTRefresh:
		return JWTRefresh
	default:
		return JWTDefault
	}
}
