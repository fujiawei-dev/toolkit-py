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

// PostEntityTemplate
// @Summary      创建
// @Description  创建
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        object  body  form.EntityTemplateCreate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template [post]
func PostEntityTemplate(router *gin.RouterGroup) {
	router.POST("/entity_template", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionCreate); !pass {
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

// PutEntityTemplate
// @Summary      修改
// @Description  修改
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int              true   "ID"
// @Param        object  body  form.EntityTemplateUpdate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template/{id} [put]
func PutEntityTemplate(router *gin.RouterGroup) {
	router.PUT("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionUpdate); !pass {
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

// DeleteEntityTemplate
// @Summary      删除
// @Description  删除
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /unit/{id} [delete]
func DeleteEntityTemplate(router *gin.RouterGroup) {
	router.DELETE("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionDelete); !pass {
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

// GetEntityTemplate
// @Summary      获取
// @Description  获取
// @Tags         实体管理
// @Accept       json
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response{result=query.EntityTemplateResult}  "操作成功"
// @Router       /entity_template/{id} [get]
func GetEntityTemplate(router *gin.RouterGroup) {
	router.GET("/entity_template/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionRead); !pass {
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

// GetEntityTemplate
// @Summary      实体列表
// @Description  获取实体列表
// @Tags         实体管理
// @Accept       json
// @Param        page        query  int     false  "页码"    default(1)
// @Param        page_size   query  int     false  "每页数量"  Enums(10, 20)  default(10)
// @Param        time_begin  query  string  false  "开始时间，比如 2021-10-01"
// @Param        time_end    query  string  false  "结束时间，比如 2022-10-01"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.EntityTemplateResult}}  "操作成功"
// @Router       /entity_templates [get]
func GetEntityTemplates(router *gin.RouterGroup) {
	router.GET("/entity_templates", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionSearch); !pass {
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
