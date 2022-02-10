{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/event"
)

var log = event.Log
var conf = config.Conf()
