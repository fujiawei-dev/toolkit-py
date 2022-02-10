{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	Debug        bool   `flag:"debug" yaml:"debug,omitempty"`
	Public       bool   `flag:"public" yaml:"public,omitempty"`
	Swagger      bool   `flag:"swagger" yaml:"swagger,omitempty"`
	ConfigFile   string `flag:"config-file" yaml:"config_file,omitempty"`
	LogLevel     string `flag:"log-level" yaml:"log_level,omitempty"`
	HttpPort     int    `flag:"http-port" yaml:"http_port,omitempty"`
	HttpMode     string `flag:"http-mode" yaml:"http_mode,omitempty"`
	DetachServer bool   `flag:"detach-server" yaml:"detach_server,omitempty"`
}

func NewFlags(flags *pflag.FlagSet) *Flags {
	f := &Flags{}

	if flags != nil {
		f.SetFlags(flags)
	}

	return f
}

// SetFlags uses options from the CLI to setup configuration overrides for the entity.
func (f *Flags) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&f.Debug, "debug", "D", false, "run in debug mode, shows additional log messages")
	flags.BoolVarP(&f.Public, "public", "P", false, "run in public mode")
	flags.BoolVarP(&f.Swagger, "swagger", "S", false, "allow access swagger docs")
	flags.StringVarP(&f.ConfigFile, "config-file", "c", "", `configuration file`)
	flags.StringVarP(&f.LogLevel, "log-level", "l", "", `logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
	flags.IntVarP(&f.HttpPort, "http-port", "p", 0, `http server port NUMBER`)
	flags.StringVarP(&f.HttpMode, "http-mode", "m", "", `debug, release or test`)
	flags.BoolVarP(&f.DetachServer, "detach-server", "d", false, "detach from the console (daemon mode)")
}
