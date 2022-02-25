{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

func SendJSON(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, query.NewResponse(http.StatusOK, nil, v))
}

func SendOK(c *gin.Context) {
	SendJSON(c, nil)
}

func SendList(c *gin.Context, list interface{}, pager form.Pager) {
	SendJSON(c, gin.H{"list": list, "pager": pager})
}

func Abort(c *gin.Context, code int) {
	resp := query.NewResponse(code, nil, nil)

	log.Error().Msgf("api: %s %s abort (%s)", c.Request.Method, c.FullPath(), resp.LowerString())

	c.JSON(query.StatusCode(code), resp)
}

func Error(c *gin.Context, code int, err error) {
	resp := query.NewResponse(code, err, nil)

	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Request.Method, c.FullPath(), resp.LowerString())
	}

	c.JSON(query.StatusCode(code), resp)
}

func ErrorInvalidParameters(c *gin.Context, err error) {
	Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c *gin.Context, err error) {
	Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c *gin.Context, err error) {
	Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c *gin.Context, err error) {
	Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c *gin.Context) {
	Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c *gin.Context) {
	Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c *gin.Context) {
	Abort(c, query.ErrInvalidPassword)
}

func ErrorExpectedOrUnexpected(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		ErrorUnexpected(c, err)
	}
}
