{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

// Resource 系统资源
type Resource string

const (
	ResourceDefault   Resource = "*"         // 默认
	ResourceConfig    Resource = "config"    // 系统运行配置
	ResourceSettings  Resource = "settings"  // 用户功能设置
	ResourceLogs      Resource = "logs"      // 日志
	ResourcePasswords Resource = "passwords" // 密码
	ResourceUsers     Resource = "users"     // 用户
	ResourceEntityTemplate    Resource = "entity_template"
)
