{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gofiber/fiber/v2"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)


func RegisterExample(router fiber.Router) {
	PostExample(router)
	PutExample(router)
	DeleteExample(router)

	GetExample(router)
	GetExamples(router)
}

func PostExample(router fiber.Router) {
	router.Post("/example", func(c *fiber.Ctx) error {
		var f form.Example

		if err := form.ShouldBind(c, &f); err != nil {
			return ErrorInvalidParams(c, err)
		}

		var m entity.Example

		if err := m.CopyFrom(f); err != nil {
			return ErrorUnexpected(c, err)
		}

		if err := m.Create(); err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendOK(c)
	})
}

func PutExample(router fiber.Router) {
	router.Put("/example", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		f := form.Pager{}

		if err := form.ShouldBind(c, &f); err != nil {
			f = form.Pager{Page: 1, PageSize: 10}
		}

		list, totalRow, err := query.Examples(f)

		if err != nil {
			return ErrorUnexpected(c, err)
		}

		f.TotalRows = totalRow

		return SendList(c, list, f)
	})
}

func DeleteExample(router fiber.Router) {
	router.Delete("/example", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		f := form.Pager{}

		if err := form.ShouldBind(c, &f); err != nil {
			f = form.Pager{Page: 1, PageSize: 10}
		}

		list, totalRow, err := query.Examples(f)

		if err != nil {
			return ErrorUnexpected(c, err)
		}

		f.TotalRows = totalRow

		return SendList(c, list, f)
	})
}

func GetExample(router fiber.Router) {
	router.Get("/example", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		f := form.Pager{}

		if err := form.ShouldBind(c, &f); err != nil {
			f = form.Pager{Page: 1, PageSize: 10}
		}

		list, totalRow, err := query.Examples(f)

		if err != nil {
			return ErrorUnexpected(c, err)
		}

		f.TotalRows = totalRow

		return SendList(c, list, f)
	})
}

func GetExamples(router fiber.Router) {
	router.Get("/examples", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		f := form.Pager{}

		if err := form.ShouldBind(c, &f); err != nil {
			f = form.Pager{Page: 1, PageSize: 10}
		}

		list, totalRow, err := query.Examples(f)

		if err != nil {
			return ErrorUnexpected(c, err)
		}

		f.TotalRows = totalRow

		return SendList(c, list, f)
	})
}
