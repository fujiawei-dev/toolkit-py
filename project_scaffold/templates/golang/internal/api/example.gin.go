{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"database/sql"
	"errors"

	"github.com/spf13/cast"
	"github.com/gin-gonic/gin"

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

func PostExample(router *gin.RouterGroup) {
	router.POST("/example", conf.JWTMiddleware(), func(c *gin.Context) {
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

func PutExample(router *gin.RouterGroup) {
	router.PUT("/example/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		id := cast.ToUint(c.Param("id"))
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

func DeleteExample(router *gin.RouterGroup) {
	router.DELETE("/example/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		id := cast.ToUint(c.Param("id"))
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

func GetExample(router *gin.RouterGroup) {
	router.GET("/example/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		id := cast.ToUint(c.Param("id"))
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

func GetExamples(router *gin.RouterGroup) {
	router.GET("/examples", conf.JWTMiddleware(), func(c *gin.Context) {
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
