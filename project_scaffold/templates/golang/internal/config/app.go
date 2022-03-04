{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"strings"
	"time"

	"{{GOLANG_MODULE}}/version"
)

type AppSetting struct {
	Description string `mapstructure:"description" yaml:"description,omitempty"`
	Title       string `mapstructure:"title" yaml:"title,omitempty"`
}

func (c *config) AppName() string {
	if version.Name != version.LibraryImport {
		return version.Name
	}

	return "program"
}

func (c *config) AppNameUpper() string {
	return strings.ToUpper(c.AppName())
}

// AppVersion returns the application version.
func (c *config) AppVersion() string {
	if version.Version != version.LibraryImport {
		return version.Version
	}

	return "1.0.0"
}

func (c *config) AppGitCommit() string {
	if version.GitCommit != version.LibraryImport {
		return version.GitCommit
	}

	return "unknown"
}

func (c *config) AppBuildTime() string {
	if version.BuildTime != version.LibraryImport {
		return version.BuildTime
	}

	return time.Now().Local().Round(time.Second).String()
}

func (c *config) Version() string {
	return fmt.Sprintf("%s, build %s at %s", conf.AppVersion(), conf.AppGitCommit(), conf.AppBuildTime())
}

func (c *config) UserAgent() string {
	return version.UserAgent(nil)
}

func (c *config) AppTitle() string {
	if c.settings.App.Title != "" {
		return c.settings.App.Title
	}

	return "Web-Application"
}

func (c *config) AppDescription() string {
	if c.settings.App.Description != "" {
		return c.settings.App.Description
	}

	return "This is a Web-Application"
}
