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

type JWTSetting struct {
	Enable     bool          `mapstructure:"enable" yaml:"enable,omitempty"`
	Expire     time.Duration `mapstructure:"expire" yaml:"expire,omitempty"`
	Field      string        `mapstructure:"field" yaml:"field,omitempty"` // Header 字段 Authorization
	Issuer     string        `mapstructure:"issuer" yaml:"issuer,omitempty"`
	Key        string        `mapstructure:"key" yaml:"key,omitempty"`
	Mode       string        `mapstructure:"mode" yaml:"mode,omitempty"` // default/refresh
	PrivateKey string        `mapstructure:"private_key" yaml:"private_key,omitempty"`
	PublicKey  string        `mapstructure:"public_key" yaml:"public_key,omitempty"`
	Scheme     string        `mapstructure:"scheme" yaml:"scheme,omitempty"` // 格式前缀

	JWTProvider
}

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
