{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/labstack/gommon/color"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/event"
)

var (
	conf *config
	once sync.Once
)

type config struct {
	gm *gorm.DB

	flags    *Flags
	settings *Settings
}

func Conf() *config {
	once.Do(func() {
		conf = &config{
			flags:    NewFlags(nil),
			settings: NewSettings(),
		}
	})

	return conf
}

func (c *config) Flags() *Flags {
	return c.flags
}

func (c *config) SetContext(ctx *cli.Context) {
	c.flags.SetContext(ctx)
}

func (c *config) Init(ctx *cli.Context)error {
	if err := c.InitSettings(ctx); err != nil {
		return err
	}

	if err := c.CreateDirectories(); err != nil {
		return err
	}

	if err := c.InitService(); err != nil {
		return err
	}

	if err := c.InitDb(); err != nil {
		return err
	}

	c.initLogger()

	c.initJWT()

	return nil
}

func (c *config) InitSettings(ctx *cli.Context) error {
	c.flags.SetContext(ctx)

	filePath := c.SettingsFile()

	log := event.Logger()

	if err := c.settings.Load(filePath); err == nil {
		log.Printf("config: settings loaded from %s", filePath)
	} else if err = c.settings.Save(filePath); err != nil {
		log.Printf("config: failed creating %s: %s", filePath, err)
	} else {
		log.Printf("config: created %s", filePath)
	}

	c.initLogger()

	return nil
}

// Settings returns the current user settings.
func (c *config) Settings() *Settings {
	return c.settings
}

// Load user settings from file.
func (s *Settings) Load(filePath string) error {
	yamlConfig, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlConfig, s)
}

// Save user settings to a file.
func (s *Settings) Save(fileName string) error {
	buf, err := yaml.Marshal(s)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, buf, os.ModePerm)
}

func (c *config) Print() {
	color.Printf("%s\n\n", color.Blue(">>> from build-time flags"))

	color.Printf("version: %s\n", c.Version())
	color.Printf("user-agent: %s\n\n", c.UserAgent())

	var content []byte

	content, _ = yaml.Marshal(c.flags)
	color.Printf("%s\n\n%s\n", color.Blue(">>> from command line flags"), content)

	content, _ = yaml.Marshal(c.settings)
	color.Printf("%s\n\n%s\n", color.Blue(">>> from "+c.SettingsFile()), content)
}
