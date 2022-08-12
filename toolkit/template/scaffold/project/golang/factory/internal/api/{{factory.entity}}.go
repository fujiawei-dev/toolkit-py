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
	AddRouteRegistrar(Post{{ factory.entity|title }})
	AddRouteRegistrar(Put{{ factory.entity|title }})
	AddRouteRegistrar(Delete{{ factory.entity|title }})

	AddRouteRegistrar(Get{{ factory.entity|title }})
	AddRouteRegistrar(Get{{ factory.entity|title }}s)
}

// Post{{ factory.entity|title }}
// @Summary      创建{{ factory.entity_chinese }}
// @Description  创建{{ factory.entity_chinese }}
// @Tags         {{ factory.entity_chinese }}管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        object  body  form.{{ factory.entity|title }}Create  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /{{ factory.entity }} [post]
func Post{{ factory.entity|title }}(router {{ web_framework_router_group }}) {
	router.{{ web_framework_post }}("/{{ factory.entity }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.Resource{{ factory.entity|title }}s, acl.ActionCreate); !pass {
			return{{ web_framework_nil }}
		}

		var f form.{{ factory.entity|title }}Create

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		var m entity.{{ factory.entity|title }}

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

// Put{{ factory.entity|title }}
// @Summary      修改{{ factory.entity_chinese }}
// @Description  修改{{ factory.entity_chinese }}
// @Tags         {{ factory.entity_chinese }}管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id      path  int                        true   "ID"
// @Param        object  body  form.{{ factory.entity|title }}Update  false  "参数"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /{{ factory.entity }}/{id} [put]
func Put{{ factory.entity|title }}(router {{ web_framework_router_group }}) {
	router.{{ web_framework_put }}("/{{ factory.entity }}/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.Resource{{ factory.entity|title }}s, acl.ActionUpdate); !pass {
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

		var m entity.{{ factory.entity|title }}
		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		// Handle null values, malicious injection, etc.
		var f form.{{ factory.entity|title }}Update

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

// Delete{{ factory.entity|title }}
// @Summary      删除{{ factory.entity_chinese }}
// @Description  删除{{ factory.entity_chinese }}
// @Tags         {{ factory.entity_chinese }}管理
// @Accept       json
// @Security     ApiKeyAuth
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response  "操作成功"
// @Router       /{{ factory.entity }}/{id} [delete]
func Delete{{ factory.entity|title }}(router {{ web_framework_router_group }}) {
	router.{{ web_framework_delete }}("/{{ factory.entity }}/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.Resource{{ factory.entity|title }}s, acl.ActionDelete); !pass {
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

		var m entity.{{ factory.entity|title }}

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

// Get{{ factory.entity|title }}
// @Summary      获取{{ factory.entity_chinese }}
// @Description  获取{{ factory.entity_chinese }}，存在必要性存疑，如前期需要再讨论
// @Tags         {{ factory.entity_chinese }}管理
// @Deprecated   true
// @Accept       json
// @Param        id  path  int  true  "ID"
// @Produce      json
// @Success      200  {object}  query.Response{result=query.{{ factory.entity|title }}Result}  "操作成功"
// @Router       /{{ factory.entity }}/{id} [get]
func Get{{ factory.entity|title }}(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/{{ factory.entity }}/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.Resource{{ factory.entity|title }}s, acl.ActionRead); !pass {
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

		var m entity.{{ factory.entity|title }}

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		var r query.{{ factory.entity|title }}Result
		m.CopyTo(&r)

		SendJSON(c, r)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}

// Get{{ factory.entity|title }}s
// @Summary      获取{{ factory.entity_chinese }}列表
// @Description  获取{{ factory.entity_chinese }}列表
// @Tags         {{ factory.entity_chinese }}管理
// @Accept       application/x-www-form-urlencoded
// @Param        page        query  int     false  "页码"                     default(1)
// @Param        page_size   query  int     false  "每页数量"                   Enums(10, 20)  default(10)
// @Param        order       query  int     false  "已定义字段排序 0 ID顺序 1 ID倒序"  Enums(0, 1)    default(1)
// @Param        time_begin  query  string  false  "开始时间前一天，比如 2021-10-01，则实际从 2021-10-02 起开始"
// @Param        time_end    query  string  false  "结束时间后一天，比如 2022-10-01，则实际到 2021-09-31 起结束"
// @Produce      json
// @Success      200  {object}  query.Response{result=Result{pager=form.Pager,list=[]query.{{ factory.entity|title }}Result}}  "操作成功"
// @Router       /{{ factory.entity }}s [get]
func Get{{ factory.entity|title }}s(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/{{ factory.entity }}s", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.Resource{{ factory.entity|title }}s, acl.ActionSearch); !pass {
			return{{ web_framework_nil }}
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		list, totalRow, err := query.{{ factory.entity|title }}s(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}
