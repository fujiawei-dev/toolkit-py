{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/kataras/iris/v12"
)

func ShouldBind(c iris.Context, ptr interface{}) (err error) {
	if err = c.ReadBody(ptr); err == nil {
		return nil
	}

	return ValidateError(err)
}
