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

	if errs, ok := err.(validator.ValidationErrors); ok {
		validatorErrors := make(ValidatorErrors, 0, len(errs))

		for _, e := range errs {
			validatorErrors = append(validatorErrors, ValidatorError{
				Key:     "error",
				Message: e.Error(),
			})
		}

		return validatorErrors
	}

	return err
}

type EchoValidator struct {
	validator *validator.Validate
}

func NewValidator() *EchoValidator {
	return &EchoValidator{validator: validator.New()}
}

func (v *EchoValidator) Validate(i interface{}) (err error) {
	return v.validator.Struct(i)
}
