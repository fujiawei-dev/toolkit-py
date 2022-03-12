{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
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

func PostExample(router *echo.Group) {
	router.POST("/example", func(c echo.Context) error {
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
	}, conf.JWTMiddleware())
}

func PutExample(router *echo.Group) {
	router.PUT("/example/:id", func(c echo.Context) error {
		id := cast.ToUint(c.Param("id"))
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
	}, conf.JWTMiddleware())
}

func DeleteExample(router *echo.Group) {
	router.DELETE("/example/:id", func(c echo.Context) error {
		id := cast.ToUint(c.Param("id"))
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
	}, conf.JWTMiddleware())
}

func GetExample(router *echo.Group) {
	router.GET("/example/:id", func(c echo.Context) error {
		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			return ErrorInvalidParameters(c, errors.New("id(uint) is required"))
		}

		var m entity.Example

		if err := m.FindByID(id); err != nil {
			return ErrorExpectedOrUnexpected(c, err)
		}

		return SendJSON(c, m)
	}, conf.JWTMiddleware())
}

func GetExamples(router *echo.Group) {
	router.GET("/examples", func(c echo.Context) error {
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
	}, conf.JWTMiddleware())
}
