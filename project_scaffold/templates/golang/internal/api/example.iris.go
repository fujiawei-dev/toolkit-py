{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"errors"

	"github.com/kataras/iris/v12"

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

func PostExample(router iris.Party) {
	router.Post("/example", func(c iris.Context) {
		var f form.ExampleCreate

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		var m entity.Example

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		m.NotNullField = sql.NullBool{Bool: true, Valid: true}

		if err := m.Create(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func PutExample(router iris.Party) {
	router.Put("/example/{id:uint}", conf.JWTMiddleware(), func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.Example
		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		// Handle null values, malicious injection, etc.
		var f form.ExampleUpdate

		if err := m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		if err := m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func DeleteExample(router iris.Party) {
	router.Delete("/example/{id:uint}", conf.JWTMiddleware(), func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.Example

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		if err := m.Delete(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func GetExample(router iris.Party) {
	router.Get("/example/{id:uint}", conf.JWTMiddleware(), func(c iris.Context) {
		id := c.Params().GetUintDefault("id", 0)
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.Example

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		SendJSON(c, m)
	})
}

func GetExamples(router iris.Party) {
	router.Get("/examples", conf.JWTMiddleware(), func(c iris.Context) {
		f := form.Pager{}

		if err := form.ShouldBind(c, &f); err != nil {
			f = form.Pager{Page: 1, PageSize: 10}
		}

		list, totalRow, err := query.Examples(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		f.TotalRows = totalRow

		SendList(c, list, f)
	})
}
