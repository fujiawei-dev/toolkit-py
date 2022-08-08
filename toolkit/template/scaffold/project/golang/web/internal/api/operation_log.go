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
	AddRouteRegistrar(DeleteOperationLog)
	AddRouteRegistrar(DeleteOperationLogs)

	AddRouteRegistrar(GetOperationLog)
	AddRouteRegistrar(GetOperationLogs)
}

func DeleteOperationLog(router {{ web_framework_router_group }}) {
	router.{{ web_framework_delete }}("/operation_log/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
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

		var m entity.OperationLog

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
	})
}

func DeleteOperationLogs(router {{ web_framework_router_group }}) {
	router.{{ web_framework_delete }}("/operation_logs", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
			return{{ web_framework_nil }}
		}

		if err := query.DeleteAllOperationLogs(); err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendOK(c)
	}{{ web_framework_jwt_down }})
}

func GetOperationLog(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/operation_log/{{ web_framework_query_id }}", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionRead); !pass {
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

		var m entity.OperationLog

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		SendJSON(c, m)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}

func GetOperationLogs(router {{ web_framework_router_group }}) {
	router.{{ web_framework_get }}("/operation_logs", {{ web_framework_jwt_up }}func(c {{ web_framework_context }}) {{ web_framework_error }}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionSearch); !pass {
			return{{ web_framework_nil }}
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{ web_framework_nil }}
		}

		list, totalRow, err := query.OperationLogs(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{ web_framework_nil }}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{ web_framework_nil }}
	}{{ web_framework_jwt_down }})
}
