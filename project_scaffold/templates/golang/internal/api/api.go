{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"errors"
	"net/http"

	"{{WEB_FRAMEWORK_IMPORT}}"
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

type JsonObject map[string]interface{}

func SendJSON(c {{WEB_CONTEXT}}, v interface{}) {
    {%- if WEB_FRAMEWORK in [".iris", ".fiber"] %}
    c.JSON(query.NewResponse(http.StatusOK, nil, v))
    {%- else %}
	c.JSON(http.StatusOK, query.NewResponse(http.StatusOK, nil, v))
    {%- endif %}
}

func SendOK(c {{WEB_CONTEXT}}) {
	SendJSON(c, nil)
}

func SendList(c {{WEB_CONTEXT}}, list interface{}, pager form.Pager) {
	SendJSON(c, JsonObject{"list": list, "pager": pager})
}

func Abort(c {{WEB_CONTEXT}}, code int) {
	resp := query.NewResponse(code, nil, nil)

    {%- if WEB_FRAMEWORK==".iris" %}
	log.Errorf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

    c.StopWithJSON(query.StatusCode(code), resp)
    {%- elif WEB_FRAMEWORK==".fiber" %}
	log.Error().Msgf("api: %s %s abort (%s)", c.Method(), c.Path(), resp.LowerString())

    c.JSON(resp)
    {%- elif WEB_FRAMEWORK==".echo" %}
	log.Errorf("api: %s %s abort (%s)", c.Request().Method, c.Path(), resp.LowerString())

	c.JSON(query.StatusCode(code), resp)
    {%- elif WEB_FRAMEWORK==".gin" %}
	log.Error().Msgf("api: %s %s abort (%s)", c.Request.Method, c.FullPath(), resp.LowerString())

	c.JSON(query.StatusCode(code), resp)
    {%- endif %}
}

func Error(c {{WEB_CONTEXT}}, code int, err error) {
	resp := query.NewResponse(code, err, nil)

    {%- if WEB_FRAMEWORK==".iris" %}
	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

    c.StopWithJSON(query.StatusCode(code), resp)
    {%- elif WEB_FRAMEWORK==".fiber" %}
	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Method(), c.Path(), resp.LowerString())
	}

	c.JSON(resp)
    {%- elif WEB_FRAMEWORK==".echo" %}
	if err != nil {
		log.Errorf("api: %s %s error (%s)", c.Request().Method, c.Path(), resp.LowerString())
	}

	c.JSON(query.StatusCode(code), resp)
    {%- elif WEB_FRAMEWORK==".gin" %}
	if err != nil {
		log.Error().Msgf("api: %s %s error (%s)", c.Request.Method, c.FullPath(), resp.LowerString())
	}

	c.JSON(query.StatusCode(code), resp)
    {%- endif %}
}

func ErrorInvalidParameters(c {{WEB_CONTEXT}}, err error) {
	Error(c, query.ErrInvalidParameters, err)
}

func ErrorUnexpected(c {{WEB_CONTEXT}}, err error) {
	Error(c, query.ErrUnexpected, err)
}

func ErrorRecordNotFound(c {{WEB_CONTEXT}}, err error) {
	Error(c, query.ErrRecordNotFound, err)
}

func ErrorRecordAlreadyExists(c {{WEB_CONTEXT}}, err error) {
	Error(c, query.ErrRecordAlreadyExists, err)
}

func AbortPermissionDenied(c {{WEB_CONTEXT}}) {
	Abort(c, query.ErrPermissionDenied)
}

func AbortUnauthorized(c {{WEB_CONTEXT}}) {
	Abort(c, http.StatusUnauthorized)
}

func AbortInvalidPassword(c {{WEB_CONTEXT}}) {
	Abort(c, query.ErrInvalidPassword)
}

func AbortNotImplemented(c {{WEB_CONTEXT}}) {
	Abort(c, http.StatusNotImplemented)
}

func ErrorExpectedOrUnexpected(c {{WEB_CONTEXT}}, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ErrorRecordNotFound(c, err)
	} else if errors.Is(err, event.ErrRecordAlreadyExists) {
		ErrorRecordAlreadyExists(c, err)
	} else if err != nil {
		ErrorUnexpected(c, err)
	}
}
