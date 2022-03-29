{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	{% set web_context="" -%}

	{% if web_framework==".gin" -%}
	"github.com/gin-gonic/gin"
	{%- set web_context="*gin.Context" -%}
	{% elif web_framework==".echo" -%}
	"github.com/labstack/echo/v4"
	{%- set web_context="echo.Context" -%}
	{% elif web_framework==".fiber" -%}
	"github.com/gofiber/fiber/v2"
	{%- set web_context="*fiber.Ctx" -%}
	{% elif web_framework==".iris" -%}
	"github.com/kataras/iris/v12"
	{%- set web_context="iris.Context" -%}
	{%- endif %}
	"gorm.io/gorm"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

{%- set return_string="" -%}
{%- set error_string="" -%}
{%- if web_framework in [".echo", ".iris"] -%}
{%- set error_string="error" -%}
{%- set return_string="return" -%}
{%- endif %}

func SendJSON(c {{web_context}}, v interface{}) {{error_string}}{
	{{return_string }}c.JSON(http.StatusOK, query.NewResponse(http.StatusOK, nil, v))
}

func SendOK(c {{web_context}}) {{error_string}}{
	{{return_string }}SendJSON(c, nil)
}

func SendList(c {{web_context}}, list interface{}, pager form.Pager) {{error_string}}{
	{{return_string }}SendJSON(c, gin.H{"list": list, "pager": pager})
}

func Abort(c {{web_context}}, code int) {{error_string}}{
	resp := query.NewResponse(code, nil, nil)

	log.Error().Msgf("api: %s %s abort (%s)", c.Request.Method, c.FullPath(), resp.LowerString())

	{{return_string }}c.JSON(query.StatusCode(code), resp)
}

func Error(c {{web_context}}, code int, err error) {{error_string}}{
	resp := query.NewResponse(code, err, nil)

	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Request.Method, c.FullPath(), resp.LowerString())
	}

	{{return_string }}c.JSON(query.StatusCode(code), resp)
}

func ErrorInvalidParameters(c {{web_context}}, err error) {{error_string}}{
	{{return_string }}Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c {{web_context}}, err error) {{error_string}}{
	{{return_string }}Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c {{web_context}}, err error) {{error_string}}{
	{{return_string }}Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c {{web_context}}, err error) {{error_string}}{
	{{return_string }}Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c {{web_context}}) {{error_string}}{
	{{return_string }}Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c {{web_context}}) {{error_string}}{
	{{return_string }}Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c {{web_context}}) {{error_string}}{
	{{return_string }}Abort(c, query.ErrInvalidPassword)
}

func ErrorExpectedOrUnexpected(c {{web_context}}, err error) {{error_string}}{
	if errors.Is(err, gorm.ErrRecordNotFound) {
		{{return_string }}ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		{{return_string }}ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		{{return_string }}ErrorUnexpected(c, err)
	}
	{%- if web_framework in [".echo", ".iris"] %}
	return nil
	{%- endif %}
	{%- if web_framework not in [".echo", ".iris"] -%}{%- endif %}
}
