{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"

	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(GetAppDescription)
}

func GetAppDescription(router fiber.Router) {
	router.Get("/app/description", func(c *fiber.Ctx) error {
		return SendJSON(c, query.GetAppDescription())
	})
}
