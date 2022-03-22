{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"{{GOLANG_MODULE}}/internal/acl"
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

func init() {
	AddRouteRegistrar(PostEntityTemplate)
	AddRouteRegistrar(PutEntityTemplate)
	AddRouteRegistrar(DeleteEntityTemplate)

	AddRouteRegistrar(GetEntityTemplate)
	AddRouteRegistrar(GetEntityTemplates)
}

func PostEntityTemplate(router *gin.RouterGroup) {
	router.POST("/entity_template", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplate, acl.ActionCreate); !pass {
			return
		}

		var f form.EntityTemplateCreate

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		var m entity.EntityTemplate

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		if err := m.Create(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func PutEntityTemplate(router *gin.RouterGroup) {
	router.PUT("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplate, acl.ActionUpdate); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.EntityTemplate
		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		// Handle null values, malicious injection, etc.
		var f form.EntityTemplateUpdate

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

func DeleteEntityTemplate(router *gin.RouterGroup) {
	router.DELETE("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplate, acl.ActionDelete); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.EntityTemplate

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

func GetEntityTemplate(router *gin.RouterGroup) {
	router.GET("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplate, acl.ActionRead); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.EntityTemplate

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		SendJSON(c, m)
	})
}

func GetEntityTemplates(router *gin.RouterGroup) {
	router.GET("/entity_templates", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplate, acl.ActionSearch); !pass {
			return
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		list, totalRow, err := query.EntityTemplates(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
	})
}
