{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"bytes"
	"os"
	"sync"

	"github.com/labstack/gommon/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

var once sync.Once

var conf *config

type config struct {
	vp *viper.Viper
	gm *gorm.DB

	flags    *Flags
	settings *Settings
}

func Conf() *config {
	once.Do(func() {
		conf = &config{
			vp:       viper.New(),
			flags:    NewFlags(nil),
			settings: NewSettings(),
		}
	})

	return conf
}

func (c *config) Flags() *Flags {
	return c.flags
}

func (c *config) SetFlags(flags *pflag.FlagSet) {
	c.flags.SetFlags(flags)
}

func (c *config) readFromViper() (err error) {
	c.vp.SetConfigFile(c.SettingsFile())

	if err = c.vp.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			return
		}

		buf, _ := yaml.Marshal(c.settings)

		if err = c.vp.ReadConfig(bytes.NewBuffer(buf)); err != nil {
			return
		}

		if err = c.vp.WriteConfig(); err != nil {
			return
		}

		return nil
	}

	return c.vp.Unmarshal(&c.settings)
}

func (c *config) InitSettings() (err error) {
	return c.readFromViper()
}

func (c *config) Init() (err error) {
	if err = c.InitSettings(); err != nil {
		return
	}

	if err = c.CreateDirectories(); err != nil {
		return
	}

	c.initLogger()

	c.initJWT()

	return c.InitDb()
}

func (c *config) Print() (err error) {
	color.Printf("%s\n\n", color.Blue(">>> from build-time flags"))

	color.Printf("version: %s\n", c.Version())
	color.Printf("user-agent: %s\n\n", c.UserAgent())

	if err = c.readFromViper(); err != nil {
		return
	}

	var content []byte

	content, _ = yaml.Marshal(c.flags)
	color.Printf("%s\n\n%s\n", color.Blue(">>> from command line flags"), content)

	content, _ = yaml.Marshal(c.settings)
	color.Printf("%s\n\n%s\n", color.Blue(">>> from "+c.SettingsFile()), content)

	return
}
