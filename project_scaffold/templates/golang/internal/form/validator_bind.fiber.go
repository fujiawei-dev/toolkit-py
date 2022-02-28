{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"
)

func ShouldBind(c *fiber.Ctx, ptr interface{}) (err error) {
	if err = c.QueryParser(ptr); err != nil {
		return err
	}

	if err = c.BodyParser(ptr); err != nil && err != fiber.ErrUnprocessableEntity {
		return err
	}

	return ValidateError(Validator().Struct(ptr))
}
