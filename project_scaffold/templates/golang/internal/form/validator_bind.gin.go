{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"
)

func ShouldBind(c *gin.Context, ptr interface{}) (err error) {
	if err = c.Bind(ptr); err == nil {
		return nil
	}

	return ValidateError(err)
}
