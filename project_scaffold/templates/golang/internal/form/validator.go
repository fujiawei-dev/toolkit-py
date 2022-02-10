{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

var v *validator.Validate
var once sync.Once

func Validator() *validator.Validate {
	once.Do(func() {
		v = validator.New()
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

func ShouldBind(c iris.Context, ptr interface{}) (err error) {
	if err = c.ReadBody(ptr); err == nil {
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
