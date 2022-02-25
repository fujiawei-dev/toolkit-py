{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ShouldBind(c echo.Context, ptr interface{}) (err error) {
	if err = c.Bind(ptr); err == nil {
		return nil
	}

	return ValidateError(err)
}

type ValidatorForEcho struct {
	validator *validator.Validate
}

func NewValidatorForEcho() *ValidatorForEcho {
	return &ValidatorForEcho{validator: Validator()}
}

func (v *ValidatorForEcho) Validate(i interface{}) (err error) {
	return v.validator.Struct(i)
}
