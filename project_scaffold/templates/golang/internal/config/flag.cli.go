{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"reflect"

	"github.com/urfave/cli"
)

const EnvVarPrefix = "{{APP_NAME_UPPER}}_"

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

// CommandFlags Command-line parameters and flags.
var CommandFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug, D",
		Usage:  "run in debug mode, shows additional log messages",
		EnvVar: EnvVarPrefix + "DEBUG",
	},
	cli.BoolFlag{
		Name:   "public, P",
		Usage:  "run in public mode",
		EnvVar: EnvVarPrefix + "PUBLIC",
	},
	cli.BoolFlag{
		Name:   "swagger, S",
		Usage:  "allow access swagger docs",
		EnvVar: EnvVarPrefix + "SWAGGER",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load initial config options from `FILENAME`",
		EnvVar: EnvVarPrefix + "CONFIG_FILE",
	},
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "http server port `NUMBER`",
		EnvVar: EnvVarPrefix + "HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		EnvVar: EnvVarPrefix + "HTTP_MODE",
	},
	cli.BoolFlag{
		Name:   "detach-server, d",
		Usage:  "detach from the console (daemon mode)",
		EnvVar: EnvVarPrefix + "DETACH_SERVER",
	},
}

func NewFlags(ctx *cli.Context) *Flags {
	f := &Flags{}
	if ctx == nil {
		return f
	}
	f.SetContext(ctx)
	return f
}

// SetContext uses options from the CLI to setup configuration overrides for the entity.
func (f *Flags) SetContext(ctx *cli.Context) {
	v := reflect.ValueOf(f).Elem()

	// Iterate through all config fields.
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)

		tagValue := v.Type().Field(i).Tag.Get("flag")

		// Automatically assign options to fields with "flag" tag.
		if tagValue != "" {
			switch fieldValue.Interface().(type) {
			case int, int64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Int64(tagValue)
					fieldValue.SetInt(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Int() == 0 {
					f := ctx.GlobalInt64(tagValue)
					fieldValue.SetInt(f)
				}
			case uint, uint64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Uint64(tagValue)
					fieldValue.SetUint(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Uint() == 0 {
					f := ctx.GlobalUint64(tagValue)
					fieldValue.SetUint(f)
				}
			case string:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.String() == "" {
					f := ctx.GlobalString(tagValue)
					fieldValue.SetString(f)
				}
			case bool:
				if ctx.IsSet(tagValue) {
					f := ctx.Bool(tagValue)
					fieldValue.SetBool(f)
				} else if ctx.GlobalIsSet(tagValue) {
					f := ctx.GlobalBool(tagValue)
					fieldValue.SetBool(f)
				}
			}
		}
	}
}
