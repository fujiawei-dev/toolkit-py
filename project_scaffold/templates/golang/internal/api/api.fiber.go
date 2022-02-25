{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func SendJSON(c *fiber.Ctx, v interface{}) error {
	return c.JSON(query.NewResponse(http.StatusOK, nil, v))
}

func SendOK(c *fiber.Ctx) error {
	return SendJSON(c, nil)
}

func SendList(c *fiber.Ctx, list interface{}, pager form.Pager) error {
	return SendJSON(c, fiber.Map{"list": list, "pager": pager})
}

func Abort(c *fiber.Ctx, code int) error {
	resp := query.NewResponse(code, nil, nil)

	log.Error().Msgf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

	return c.JSON(resp)
}

func Error(c *fiber.Ctx, code int, err error) error {
	resp := query.NewResponse(code, err, nil)

	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

	return c.JSON(resp)
}

func ErrorInvalidParameters(c *fiber.Ctx, err error) error {
	return Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c *fiber.Ctx, err error) error {
	return Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c *fiber.Ctx, err error) error {
	return Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c *fiber.Ctx, err error) error {
	return Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c *fiber.Ctx) error {
	return Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c *fiber.Ctx) error {
	return Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c *fiber.Ctx) error {
	return Abort(c, query.ErrInvalidPassword)
}

func ErrorExpectedOrUnexpected(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		return ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		return ErrorUnexpected(c, err)
	}

	return nil
}
