{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"{{GOLANG_MODULE}}/pkg/fs"
)

type ServiceSetting struct {
	RemoteUrl string `mapstructure:"remote_url" yaml:"remote_url,omitempty"`

	codeRegion    map[string]string
	codeLocations map[string][2]string
}

func (c *config) RemoteUrl() string {
	if c.settings.Service.RemoteUrl != "" {
		return c.settings.Service.RemoteUrl
	}

	return "http://localhost"
}

func (c *config) InitService() error {
	{
		file := fs.Join(c.AssetsPath(), "geography/CodeRegion.json")
		if !fs.Exists(file) {
			panic(fmt.Sprintf("%s not found", fs.MustAbs(file)))
		}

		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(buf, &c.settings.Service.codeRegion)
		if err != nil {
			return err
		}
	}

	{
		file := fs.Join(c.AssetsPath(), "geography/CodeLocation.json")
		if !fs.Exists(file) {
			panic(fmt.Sprintf("%s not found", fs.MustAbs(file)))
		}

		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(buf, &c.settings.Service.codeLocations)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *config) GetRegionByCode(code string) (string, [2]string) {
	return c.settings.Service.codeRegion[code], c.settings.Service.codeLocations[code]
}
