{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"
)

func ShouldBind(c *fiber.Ctx, ptr interface{}) (err error) {
	if err = c.BodyParser(ptr); err != nil {
		return err
	}

	return ValidateError(Validator().Struct(ptr))
}
