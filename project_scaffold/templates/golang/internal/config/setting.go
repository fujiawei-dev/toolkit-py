{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"time"

	"gorm.io/gorm/logger"
)

type DatabaseSetting struct {
	DbName   string          `mapstructure:"db_name" yaml:"db_name,omitempty"`
	DbPath   string          `mapstructure:"db_path" yaml:"db_path,omitempty"`
	DbType   string          `mapstructure:"db_type" yaml:"db_type,omitempty"`
	HostPort string          `mapstructure:"host_port" yaml:"host_port,omitempty"`
	LogLevel logger.LogLevel `mapstructure:"log_level" yaml:"log_level,omitempty"`
	Password string          `mapstructure:"password" yaml:"password,omitempty"`
	Username string          `mapstructure:"username" yaml:"username,omitempty"`
}

type StaticSetting struct {
	BackupPath  string `mapstructure:"backup_path" yaml:"backup_path,omitempty"`
	StoragePath string `mapstructure:"storage_path" yaml:"storage_path,omitempty"`
	AssetsPath  string `mapstructure:"assets_path" yaml:"assets_path,omitempty"`
}

type Settings struct {
	App      AppSetting      `mapstructure:"app" yaml:"app,omitempty"`
	Database DatabaseSetting `mapstructure:"database" yaml:"database,omitempty"`
	JWT      JWTSetting      `mapstructure:"jwt" yaml:"jwt,omitempty"`
	Logger   LoggerSetting   `mapstructure:"logger" yaml:"logger,omitempty"`
	Server   ServerSetting   `mapstructure:"server" yaml:"server,omitempty"`
	Service  ServiceSetting  `mapstructure:"service" yaml:"service,omitempty"`
	Static   StaticSetting   `mapstructure:"static" yaml:"static,omitempty"`
}

func NewSettings() *Settings {
	return &Settings{
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
		Server: ServerSetting{
			BasePath: "/api/v1",
			HttpHost: "localhost",
			HttpPort: 8787,
			RunMode:  "debug",
		},
	}
}
