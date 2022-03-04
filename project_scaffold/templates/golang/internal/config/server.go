{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"runtime"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
)

type ServerSetting struct {
	BasePath string `mapstructure:"base_path" yaml:"base_path,omitempty"`
	HttpHost string `mapstructure:"http_host" yaml:"http_host,omitempty"`
	HttpPort int    `mapstructure:"http_port" yaml:"http_port,omitempty"`
	RunMode  string `mapstructure:"run_mode" yaml:"run_mode,omitempty"`
}

// HttpMode returns the server mode.
func (c *config) HttpMode() string {
	if c.flags.HttpMode != "" {
		return c.flags.HttpMode
	}

	if c.settings.Server.RunMode != "" {
		return c.settings.Server.RunMode
	}

	return ReleaseMode
}

func (c *config) Debug() bool {
	return c.flags.Debug
}

func (c *config) Public() bool {
	return c.flags.Public
}

func (c *config) Swagger() bool {
	return c.flags.Swagger
}

// DetachServer tests if server should detach from console (daemon mode).
func (c *config) DetachServer() bool {
	if runtime.GOOS == "windows" {
		return false
	}

	return c.flags.DetachServer
}

// HttpHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *config) HttpHost() string {
	if c.settings.Server.HttpHost != "" {
		return c.settings.Server.HttpHost
	}

	return "localhost"
}

// HttpPort returns the built-in HTTP server port.
func (c *config) HttpPort() int {
	if c.flags.HttpPort != 0 {
		return c.flags.HttpPort
	}

	if c.settings.Server.HttpPort != 0 {
		return c.settings.Server.HttpPort
	}

	return 10080
}

func (c *config) HttpHostPort() string {
	return fmt.Sprintf("%s:%d", c.HttpHost(), c.HttpPort())
}

func (c *config) InternalHttpHostPort() string {
	if c.Public() {
		return fmt.Sprintf("0.0.0.0:%d", c.HttpPort())
	}

	return c.HttpHostPort()
}

func (c *config) ExternalHttpHostPort() string {
	if !c.Public() || c.HttpHost() == "0.0.0.0" {
		return fmt.Sprintf("localhost:%d", c.HttpPort())
	}

	return c.HttpHostPort()
}

func (c *config) BasePath() string {
	if c.settings.Server.BasePath != "" {
		return c.settings.Server.BasePath
	}

	return "/api/v1"
}
