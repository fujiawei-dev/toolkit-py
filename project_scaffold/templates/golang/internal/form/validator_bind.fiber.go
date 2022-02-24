{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ShouldBind(c *fiber.Ctx, ptr interface{}) (err error) {
	if err = c.BodyParser(ptr); err == nil {
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
