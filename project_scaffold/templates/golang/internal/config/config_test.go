{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

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
