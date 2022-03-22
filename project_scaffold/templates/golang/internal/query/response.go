{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))), nil
}

// Response represents an error that occurred while handling a request.
type Response struct {
	Code   int         `json:"code" example:"200"`             // 错误码/状态码，正常请求无错误的情况，值为 0 或者 200
	Msg    string      `json:"message,omitempty" example:"OK"` // 错误码/状态码的文本描述
	Err    string      `json:"error,omitempty" example:"错误详情，仅测试模式可见，前端可忽略该字段"`
	Result interface{} `json:"result,omitempty"` // 请求结果数据
}

func NewResponse(code int, err error, result interface{}) Response {
	r := Response{Code: code, Msg: Message(code), Result: result}

	if conf.Debug() && err != nil {
		r.Err = err.Error()
	}

	return r
}

func (r Response) String() string {
	if r.Err != "" {
		return r.Err
	}

	return r.Msg
}

func (r Response) LowerString() string {
	return strings.ToLower(r.String())
}

func (r Response) Error() string {
	return r.Err
}

func (r Response) Success() bool {
	return r.Err == "" && r.Code < 400
}

func Message(code int) (message string) {
	message = Messages[code]

	if message != "" {
		return message
	}

	return http.StatusText(code)
}

func StatusCode(code int) int {
	if code < 1000 {
		return code
	}

	return code / 100
}

const (
	// 命名方式：基础状态码 * 100，然后枚举

	// 参数类错误
	ErrInvalidParameters   = 40001
	ErrRecordNotFound      = 40401
	ErrRecordAlreadyExists = 40901

	// 鉴权类错误
	ErrInvalidPassword  = 40101
	ErrPermissionDenied = 40301

	// 未知类错误
	ErrUnexpected = 50001
)

var Messages = map[int]string{
	ErrInvalidParameters:   "Invalid parameters, please try again later",
	ErrRecordNotFound:      "Record not found",
	ErrRecordAlreadyExists: "Record already exists",

	ErrInvalidPassword:  "Invalid password, please try again",
	ErrPermissionDenied: "Don't have permission",

	ErrUnexpected: "Unexpected error, please try again",
}
