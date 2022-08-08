package api

import (
	"errors"
	"net/http"

	"{{ web_framework_import }}"
	"gorm.io/gorm"

	"{{ main_module }}/internal/config"
	"{{ main_module }}/internal/event"
	"{{ main_module }}/internal/form"
	"{{ main_module }}/internal/query"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

type JsonObject map[string]interface{}

func SendJSON(c {{ web_framework_context }}, v interface{}) {
    {%- if web_framework in [".iris", ".fiber"] %}
    c.JSON(query.NewResponse(http.StatusOK, nil, v))
    {%- else %}
	c.JSON(http.StatusOK, query.NewResponse(http.StatusOK, nil, v))
    {%- endif %}
}

func SendOK(c {{ web_framework_context }}) {
	SendJSON(c, nil)
}

func SendList(c {{ web_framework_context }}, list interface{}, pager form.Pager) {
	SendJSON(c, JsonObject{"list": list, "pager": pager})
}

func Abort(c {{ web_framework_context }}, code int) {
	resp := query.NewResponse(code, nil, nil)

    {%- if web_framework==".iris" %}
	log.Errorf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

    c.StopWithJSON(query.StatusCode(code), resp)
    {%- elif web_framework==".fiber" %}
	log.Error().Msgf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

    c.JSON(resp)
    {%- elif web_framework==".echo" %}
	log.Errorf("api: %s %s abort (%s)", c.Request().Method, c.Path(), resp.LowerString())

	c.JSON(query.StatusCode(code), resp)
    {%- elif web_framework==".gin" %}
	log.Error().Msgf("api: %s %s abort (%s)", c.Request.Method, c.FullPath(), resp.LowerString())

	c.JSON(query.StatusCode(code), resp)
    {%- endif %}
}

func Error(c {{ web_framework_context }}, code int, err error) {
	resp := query.NewResponse(code, err, nil)

    {%- if web_framework==".iris" %}
	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

    c.StopWithJSON(query.StatusCode(code), resp)
    {%- elif web_framework==".fiber" %}
	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

	c.JSON(resp)
    {%- elif web_framework==".echo" %}
	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Request().Method, c.Path(), resp.LowerString())
	}

	c.JSON(query.StatusCode(code), resp)
    {%- elif web_framework==".gin" %}
	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Request.Method, c.FullPath(), resp.LowerString())
	}

	c.JSON(query.StatusCode(code), resp)
    {%- endif %}
}

func ErrorInvalidParameters(c {{ web_framework_context }}, err error) {
	Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c {{ web_framework_context }}, err error) {
	Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c {{ web_framework_context }}, err error) {
	Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c {{ web_framework_context }}, err error) {
	Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c {{ web_framework_context }}) {
	Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c {{ web_framework_context }}) {
	Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c {{ web_framework_context }}) {
	Abort(c, query.ErrInvalidPassword)
}

func AbortNotImplemented(c {{ web_framework_context }}) {
	Abort(c, http.StatusNotImplemented)
}

func ErrorExpectedOrUnexpected(c {{ web_framework_context }}, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		ErrorUnexpected(c, err)
	}
}
