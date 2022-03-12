{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"

	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(PostExample)
	AddRouteRegistrar(PutExample)
	AddRouteRegistrar(DeleteExample)

	AddRouteRegistrar(GetExample)
	AddRouteRegistrar(GetExamples)
}

func PostExample(router fiber.Router) {
	router.Post("/example", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		var f form.ExampleCreate

		if err := form.ShouldBind(c, &f); err != nil {
			return ErrorInvalidParameters(c, err)
		}

		var m entity.Example

		if err := m.CopyFrom(f); err != nil {
			return ErrorUnexpected(c, err)
		}

		m.NotNullField = sql.NullBool{Bool: true, Valid: true}

		if err := m.Create(); err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendOK(c)
	})
}

func PutExample(router fiber.Router) {
	router.Put("/example/:id", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		id := cast.ToUint(c.Params("id"))
		if id == 0 {
			return ErrorInvalidParameters(c, errors.New("id(uint) is required"))
		}

		var m entity.Example
		if err := m.FindByID(id); err != nil {
			return ErrorExpectedOrUnexpected(c, err)
		}

		// Handle null values, malicious injection, etc.
		var f form.ExampleUpdate

		if err := m.CopyTo(&f); err != nil {
			return ErrorUnexpected(c, err)
		}

		if err := form.ShouldBind(c, &f); err != nil {
			return ErrorInvalidParameters(c, err)
		}

		if err := m.CopyFrom(f); err != nil {
			return ErrorUnexpected(c, err)
		}

		if err := m.Save(); err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendOK(c)
	})
}

func DeleteExample(router fiber.Router) {
	router.Delete("/example/:id", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		id := cast.ToUint(c.Params("id"))
		if id == 0 {
			return ErrorInvalidParameters(c, errors.New("id(uint) is required"))
		}

		var m entity.Example

		if err := m.FindByID(id); err != nil {
			return ErrorExpectedOrUnexpected(c, err)
		}

		if err := m.Delete(); err != nil {
			return ErrorUnexpected(c, err)
		}

		return SendOK(c)
	})
}

func GetExample(router fiber.Router) {
	router.Get("/example/:id", conf.JWTMiddleware(), func(c *fiber.Ctx) error {
		id := cast.ToUint(c.Params("id"))
		if id == 0 {
			return ErrorInvalidParameters(c, errors.New("id(uint) is required"))
		}

		var m entity.Example

		if err := m.FindByID(id); err != nil {
			return ErrorExpectedOrUnexpected(c, err)
		}

		return SendJSON(c, m)
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
