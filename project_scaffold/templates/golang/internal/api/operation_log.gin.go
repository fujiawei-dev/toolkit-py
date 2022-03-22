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
	AddRouteRegistrar(DeleteOperationLog)
	AddRouteRegistrar(DeleteOperationLogs)

	AddRouteRegistrar(GetOperationLog)
	AddRouteRegistrar(GetOperationLogs)
}

func DeleteOperationLog(router *gin.RouterGroup) {
	router.DELETE("/operation_log/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.OperationLog

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

func DeleteOperationLogs(router *gin.RouterGroup) {
	router.DELETE("/operation_logs", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionDelete); !pass {
			return
		}

		if err := query.DeleteAllOperationLogs(); err != nil {
			ErrorUnexpected(c, err)
			return
		}

		SendOK(c)
	})
}

func GetOperationLog(router *gin.RouterGroup) {
	router.GET("/operation_log/:id", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionRead); !pass {
			return
		}

		id := cast.ToUint(c.Param("id"))
		if id == 0 {
			ErrorInvalidParameters(c, errors.New("id(uint) is required"))
			return
		}

		var m entity.OperationLog

		if err := m.FindByID(id); err != nil {
			ErrorExpectedOrUnexpected(c, err)
			return
		}

		SendJSON(c, m)
	})
}

func GetOperationLogs(router *gin.RouterGroup) {
	router.GET("/operation_logs", conf.JWTMiddleware(), func(c *gin.Context) {
		if _, pass := Auth(c, acl.ResourceLogs, acl.ActionSearch); !pass {
			return
		}

		f := form.SearchPager{}

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return
		}

		list, totalRow, err := query.OperationLogs(f)

		if err != nil {
			ErrorUnexpected(c, err)
			return
		}

		f.TotalRows = totalRow

		SendList(c, list, f.Pager)
	})
}
