{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
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

func SendJSON(c iris.Context, v interface{}) {
	_, _ = c.JSON(query.NewResponse(http.StatusOK, nil, v))
}

func SendOK(c iris.Context) {
	SendJSON(c, nil)
}

func SendList(c iris.Context, list interface{}, pager form.Pager) {
	SendJSON(c, iris.Map{"list": list, "pager": pager})
}

func Abort(c iris.Context, code int) {
	resp := query.NewResponse(code, nil, nil)

	log.Errorf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

	c.StopWithJSON(query.StatusCode(code), resp)
}

func Error(c iris.Context, code int, err error) {
	resp := query.NewResponse(code, err, nil)

	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

	c.StopWithJSON(query.StatusCode(code), resp)
}

func ErrorInvalidParameters(c iris.Context, err error) {
	Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c iris.Context, err error) {
	Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c iris.Context, err error) {
	Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c iris.Context, err error) {
	Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c iris.Context) {
	Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c iris.Context) {
	Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c iris.Context) {
	Abort(c, query.ErrInvalidPassword)
}

func ErrorExpectedOrUnexpected(c iris.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		ErrorUnexpected(c, err)
	}
}
