package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	if err := Conf().Init(); err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", conf.settings)
	}
}
