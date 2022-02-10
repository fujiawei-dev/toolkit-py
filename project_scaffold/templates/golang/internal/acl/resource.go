{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type Resource string

const (
	ResourceDefault       Resource = "*"
	ResourceConfig        Resource = "config"
	ResourceConfigOptions Resource = "config_options"
	ResourceSettings      Resource = "settings"
	ResourceLogs          Resource = "logs"
	ResourcePasswords     Resource = "passwords"
	ResourceUsers         Resource = "users"
)
