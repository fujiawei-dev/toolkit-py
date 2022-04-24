{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"

	"{{WEB_FRAMEWORK_IMPORT}}"
    {%- if WEB_FRAMEWORK not in [".iris"] %}
	"github.com/spf13/cast"
    {%- endif %}

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
// @Summary      创建实体
// @Description  创建实体
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        object  body  form.EntityTemplateCreate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template [post]
func PostEntityTemplate(router {{ROUTER_GROUP}}) {
	router.{{POST_STRING}}("/entity_template", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionCreate); !pass {
			return{{NIL_STRING}}
		}

		var f form.EntityTemplateCreate

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		var m entity.EntityTemplate

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err := m.Create(); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	})
}

// PutEntityTemplate
// @Summary      修改实体
// @Description  修改实体
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int              true   "ID"
// @Param        object  body  form.EntityTemplateUpdate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template/{id} [put]
func PutEntityTemplate(router {{ROUTER_GROUP}}) {
	router.{{PUT_STRING}}("/entity_template/{{QUERY_ID}}", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionUpdate); !pass {
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{NIL_STRING}}
		}

		var m entity.EntityTemplate
		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		// Handle null values, malicious injection, etc.
		var f form.EntityTemplateUpdate

		if err := m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err := m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

// DeleteEntityTemplate
// @Summary      删除实体
// @Description  删除实体
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template/{id} [delete]
func DeleteEntityTemplate(router {{ROUTER_GROUP}}) {
	router.{{DELETE_STRING}}("/entity_template/{{QUERY_ID}}", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionDelete); !pass {
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{NIL_STRING}}
		}

		var m entity.EntityTemplate

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		if err := m.Delete(); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

// GetEntityTemplate
// @Summary      获取实体
// @Description  获取实体，存在必要性存疑，如前期需要再讨论
// @Tags         实体管理
// @Deprecated   true
// @Accept       json
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response{result=query.EntityTemplateResult}  "操作成功"
// @Router       /entity_template/{id} [get]
func GetEntityTemplate(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/entity_template/{{QUERY_ID}}", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionRead); !pass {
			return{{NIL_STRING}}
		}

        {% if WEB_FRAMEWORK==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif WEB_FRAMEWORK==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif WEB_FRAMEWORK==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif WEB_FRAMEWORK==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{NIL_STRING}}
		}

		var m entity.EntityTemplate

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		var r query.EntityTemplateResult
		m.CopyTo(&r)

		SendJSON(c, r)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

// GetEntityTemplates
// @Summary      获取实体列表
// @Description  获取实体列表
// @Tags         实体管理
// @Accept       application/x-www-form-urlencoded
// @Param        page        query  int     false  "页码"    default(1)
// @Param        page_size   query  int     false  "每页数量"  Enums(10, 20)  default(10)
// @Param        time_begin  query  string  false  "开始时间前一天，比如 2021-10-01，则实际从 2021-10-02 起开始"
// @Param        time_end    query  string  false  "结束时间后一天，比如 2022-10-01，则实际到 2021-09-31 起结束"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.EntityTemplateResult}}  "操作成功"
// @Router       /entity_templates [get]
func GetEntityTemplates(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/entity_templates", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionSearch); !pass {
			return{{NIL_STRING}}
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		list, totalRow, err := query.EntityTemplates(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}
