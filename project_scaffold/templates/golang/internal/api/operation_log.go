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
	AddRouteRegistrar(DeleteOperationLog)
	AddRouteRegistrar(DeleteOperationLogs)

	AddRouteRegistrar(GetOperationLog)
	AddRouteRegistrar(GetOperationLogs)
}

func DeleteOperationLog(router {{ROUTER_GROUP}}) {
	router.{{DELETE_STRING}}("/operation_log/{{QUERY_ID}}", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
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

		var m entity.OperationLog

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
	})
}

func DeleteOperationLogs(router {{ROUTER_GROUP}}) {
	router.{{DELETE_STRING}}("/operation_logs", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
			return{{NIL_STRING}}
		}

		if err := query.DeleteAllOperationLogs(); err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendOK(c)
	}{{WEB_JWT_DOWN}})
}

func GetOperationLog(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/operation_log/:id", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionRead); !pass {
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

		var m entity.OperationLog

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return{{NIL_STRING}}
		}

		SendJSON(c, m)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}

func GetOperationLogs(router {{ROUTER_GROUP}}) {
	router.{{GET_STRING}}("/operation_logs", {{WEB_JWT_UP}}func(c {{WEB_CONTEXT}}) {{ERROR_STRING}}{
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionSearch); !pass {
			return{{NIL_STRING}}
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		list, totalRow, err := query.OperationLogs(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return{{NIL_STRING}}
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
		return{{NIL_STRING}}
	}{{WEB_JWT_DOWN}})
}
