package api

import (
	"errors"

	"{{ web_framework_import }}"
    {%- if web_framework not in [".iris"] %}
	"github.com/spf13/cast"
    {%- endif %}

	"{{ main_module }}/internal/acl"
	"{{ main_module }}/internal/entity"
	"{{ main_module }}/internal/form"
	"{{ main_module }}/internal/query"
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
func PostEntityTemplate(router {{ web_framework_router_group }}) {
	router.{{ web_framework_post }}("/entity_template", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionCreate); !pass {
			return{{ web_framework_nil }}
		}

		var f form.EntityTemplateCreate

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		var m entity.EntityTemplate

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err := m.Create(); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	})
}

// PutEntityTemplate
// @Summary      修改实体
// @Description  修改实体
// @Tags         实体管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int                        true   "ID"
// @Param        object  body  form.EntityTemplateUpdate  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /entity_template/{id} [put]
func PutEntityTemplate(router {{ web_framework_router_group }}) {
	router.{{ web_framework_put }}("/entity_template/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionUpdate); !pass {
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif web_framework==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif web_framework==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif web_framework==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{ web_framework_nil }}
		}

		var m entity.EntityTemplate
		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		// Handle null values, malicious injection, etc.
		var f form.EntityTemplateUpdate

		if err := m.CopyTo(&f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		if err := m.CopyFrom(f); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err := m.Save(); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
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
func DeleteEntityTemplate(router {{ web_framework_router_group }}) {
	router.{{ web_framework_delete }}("/entity_template/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionDelete); !pass {
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif web_framework==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif web_framework==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif web_framework==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{ web_framework_nil }}
		}

		var m entity.EntityTemplate

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		if err := m.Delete(); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
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
func GetEntityTemplate(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/entity_template/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionRead); !pass {
			return{{ web_framework_nil }}
		}

        {% if web_framework==".iris" -%}
		id := c.Params().GetUintDefault("id", 0)
        {% elif web_framework==".fiber" -%}
		id := cast.ToUint(c.Params("id"))
        {% elif web_framework==".echo" -%}
		id := cast.ToUint(c.Param("id"))
        {% elif web_framework==".gin" -%}
		id := cast.ToUint(c.Param("id"))
        {% endif -%}
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return{{ web_framework_nil }}
		}

		var m entity.EntityTemplate

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		var r query.EntityTemplateResult
		m.CopyTo(&r)

		SendJSON(c, r)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}

// GetEntityTemplates
// @Summary      获取实体列表
// @Description  获取实体列表
// @Tags         实体管理
// @Accept       application/x-www-form-urlencoded
// @Param        page        query  int     false  "页码"                     default(1)
// @Param        page_size   query  int     false  "每页数量"                   Enums(10, 20)  default(10)
// @Param        order       query  int     false  "已定义字段排序 0 ID顺序 1 ID倒序"  Enums(0, 1)    default(1)
// @Param        time_begin  query  string  false  "开始时间前一天，比如 2021-10-01，则实际从 2021-10-02 起开始"
// @Param        time_end    query  string  false  "结束时间后一天，比如 2022-10-01，则实际到 2021-09-31 起结束"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.EntityTemplateResult}}  "操作成功"
// @Router       /entity_templates [get]
func GetEntityTemplates(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/entity_templates", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceEntityTemplates, acl.ActionSearch); !pass {
			return{{ web_framework_nil }}
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		list, totalRow, err := query.EntityTemplates(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}
