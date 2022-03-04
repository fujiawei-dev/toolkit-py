{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

type ServiceSetting struct {
	RemoteUrl string `mapstructure:"remote_url" yaml:"remote_url,omitempty"`
}

func (c *config) RemoteUrl() string {
	if c.settings.Service.RemoteUrl != "" {
		return c.settings.Service.RemoteUrl
	}

	return "http://localhost"
}
