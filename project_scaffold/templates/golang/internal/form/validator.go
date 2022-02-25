{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	v     *validator.Validate
	vOnce sync.Once
)

func Validator() *validator.Validate {
	vOnce.Do(func() {
		v = validator.New()
		v.SetTagName("binding")
	})

	return v
}

type ValidatorError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func (v *ValidatorError) Error() string {
	return v.Message
}

type ValidatorErrors []ValidatorError

func (v ValidatorErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v ValidatorErrors) Error() string {
	return strings.Join(v.Errors(), ", ")
}

func ValidateError(err error) error {
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
