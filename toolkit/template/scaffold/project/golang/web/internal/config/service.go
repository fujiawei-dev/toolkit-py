package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"{{ main_module }}/pkg/fs"
)

type ServiceSetting struct {
	RemoteUrl string `mapstructure:"remote_url" yaml:"remote_url,omitempty"`

	codeRegion               map[string]string
	codeLocations            map[string][2]string
	codeRegionParentChildren map[string][]string
}

func (c *config) RemoteUrl() string {
	if c.settings.Service.RemoteUrl != "" {
		return c.settings.Service.RemoteUrl
	}

	return "http://127.0.0.1"
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

	{
		file := fs.Join(c.AssetsPath(), "geography/CodeRegionParentChildren.json")
		if !fs.Exists(file) {
			panic(fmt.Sprintf("%s not found", fs.MustAbs(file)))
		}

		buf, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal(buf, &c.settings.Service.codeRegionParentChildren)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *config) GetRegionByCode(code string) (string, [2]string) {
	return c.settings.Service.codeRegion[code], c.settings.Service.codeLocations[code]
}

func (c *config) GetChildrenByParentCode(code string) []string {
	return c.settings.Service.codeRegionParentChildren[code]
}
