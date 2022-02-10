{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"crypto/rsa"
	"sync"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
	"gorm.io/gorm/logger"
)

type AppSetting struct {
	Title       string `mapstructure:"title" yaml:"title,omitempty"`
	Description string `mapstructure:"description" yaml:"description,omitempty"`
}

type ServerSetting struct {
	BasePath string `mapstructure:"base_path" yaml:"base_path,omitempty"`
	HttpHost string `mapstructure:"http_host" yaml:"http_host,omitempty"`
	HttpPort int    `mapstructure:"http_port" yaml:"http_port,omitempty"`
	RunMode  string `mapstructure:"run_mode" yaml:"run_mode,omitempty"`
}

type JWTSetting struct {
	Enable     bool          `mapstructure:"enable" yaml:"enable,omitempty"`
	Key        string        `mapstructure:"key" yaml:"key,omitempty"`
	Issuer     string        `mapstructure:"issuer" yaml:"issuer,omitempty"`
	Scheme     string        `mapstructure:"scheme" yaml:"scheme,omitempty"` // 格式前缀
	Field      string        `mapstructure:"field" yaml:"field,omitempty"`   // Header 字段 Authorization
	Expire     time.Duration `mapstructure:"expire" yaml:"expire,omitempty"`
	Mode       string        `mapstructure:"mode" yaml:"mode,omitempty"` // default/refresh
	PrivateKey string        `mapstructure:"private_key" yaml:"private_key,omitempty"`
	PublicKey  string        `mapstructure:"public_key" yaml:"public_key,omitempty"`

	once         sync.Once
	signatureAlg jwt.Alg
	signer       *jwt.Signer
	verifier     *jwt.Verifier
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
}

type DatabaseSetting struct {
	DbType   string          `mapstructure:"db_type" yaml:"db_type,omitempty"`
	DbPath   string          `mapstructure:"db_path" yaml:"db_path,omitempty"`
	DbName   string          `mapstructure:"db_name" yaml:"db_name,omitempty"`
	Username string          `mapstructure:"username" yaml:"username,omitempty"`
	Password string          `mapstructure:"password" yaml:"password,omitempty"`
	HostPort string          `mapstructure:"host_port" yaml:"host_port,omitempty"`
	LogLevel logger.LogLevel `mapstructure:"log_level" yaml:"log_level,omitempty"`
}

type LoggerSetting struct {
	Filename   string `mapstructure:"filename" yaml:"filename,omitempty"`       // 文件名
	SavePath   string `mapstructure:"save_path" yaml:"save_path,omitempty"`     // 保存路径
	Level      string `mapstructure:"level" yaml:"level,omitempty"`             // 日志记录级别
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size,omitempty"`       // 文件大小限制
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age,omitempty"`         // 文件保留时间限制
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups,omitempty"` // 文件保留数量限制
}

type StaticSetting struct {
	StoragePath string `mapstructure:"storage_path" yaml:"storage_path,omitempty"`
	BackupPath  string `mapstructure:"backup_path" yaml:"backup_path,omitempty"`
}

type Settings struct {
	App      AppSetting      `mapstructure:"app" yaml:"app,omitempty"`
	Server   ServerSetting   `mapstructure:"server" yaml:"server,omitempty"`
	JWT      JWTSetting      `mapstructure:"jwt" yaml:"jwt,omitempty"`
	Database DatabaseSetting `mapstructure:"database" yaml:"database,omitempty"`
	Logger   LoggerSetting   `mapstructure:"logger" yaml:"logger,omitempty"`
	Static   StaticSetting   `mapstructure:"static" yaml:"static,omitempty"`
}

func NewSettings() *Settings {
	return &Settings{
		Server: ServerSetting{
			BasePath: "/api/v1",
			HttpHost: "localhost",
			HttpPort: 20080,
			RunMode:  "debug",
		},
		JWT: JWTSetting{
			Enable: true,
			Expire: 3 * time.Hour,
		},
		Logger: LoggerSetting{
			Level:      "debug",
			MaxSize:    300,
			MaxAge:     30,
			MaxBackups: 15,
		},
	}
}
