{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
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

func SendJSON(c echo.Context, v interface{}) error {
	return c.JSON(http.StatusOK, query.NewResponse(http.StatusOK, nil, v))
}

func SendOK(c echo.Context) error {
	return SendJSON(c, nil)
}

func SendList(c echo.Context, list interface{}, pager form.Pager) error {
	return SendJSON(c, echo.Map{"list": list, "pager": pager})
}

func Abort(c echo.Context, code int) error {
	resp := query.NewResponse(code, nil, nil)

	log.Errorf("api: %s %s abort (%s)", c.Request().Method, c.Path(), resp.LowerString())

	return c.JSON(query.StatusCode(code), resp)
}

func Error(c echo.Context, code int, err error) error {
	resp := query.NewResponse(code, err, nil)

	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Request().Method, c.Path(), resp.LowerString())
	}

	return c.JSON(query.StatusCode(code), resp)
}

func ErrorInvalidParameters(c echo.Context, err error) error {
	return Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c echo.Context, err error) error {
	return Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c echo.Context, err error) error {
	return Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c echo.Context, err error) error {
	return Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c echo.Context) error {
	return Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c echo.Context) error {
	return Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c echo.Context) error {
	return Abort(c, query.ErrInvalidPassword)
}

func ErrorExpectedOrUnexpected(c echo.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		return ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		return ErrorUnexpected(c, err)
	}

	return nil
}
