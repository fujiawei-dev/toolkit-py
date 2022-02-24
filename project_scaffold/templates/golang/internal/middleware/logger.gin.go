{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	// Optional. Default value is gin.defaultLogFormatter
	Formatter LogFormatter

	// DisableColor disables the colors and styles.
	DisableColor bool

	// Output is a writer where logs are written.
	// Optional. Default value is gin.DefaultWriter.
	Output io.Writer

	// SkipPaths is a url path array which logs are not written.
	// Optional.
	SkipPaths []string
}

// LogFormatter gives the signature of the formatter function passed to LoggerWithFormatter
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	Request *http.Request

	// DisableColor disables the colors and styles.
	DisableColor bool
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// BodySize is the size of the Response Body
	BodySize int
	// Keys are the keys set on the request's context.
	Keys map[string]interface{}
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (p *LogFormatterParams) MethodColor() string {
	method := p.Method

	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string

	formatter := `{"time": "%v", "status": "%s%d%s", "latency": "%v", ` +
		`"client_ip": "%s", "method": "%s%s%s", "path": "%s", "error": "%s"}` + "\n"

	if !param.DisableColor {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
		formatter = "[GIN] %v |%s %d %s| %v | %s |%s %s %s %#v %s\n"
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf(
		formatter,
		param.TimeStamp.Format("2006.01.02 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// ErrorLogger returns a handlerfunc for any error type.
func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

// ErrorLoggerT returns a handlerfunc for a given error type.
func ErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default, gin.DefaultWriter = os.Stdout.
func Logger() gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter instance a Logger middleware with the specified log format function.
func LoggerWithFormatter(f LogFormatter) gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: f,
	})
}

// LoggerWithWriter instance a Logger middleware with the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func LoggerWithWriter(out io.Writer, skipPaths ...string) gin.HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Output:    out,
		SkipPaths: skipPaths,
	})
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	skipPaths := conf.SkipPaths

	var skip map[string]struct{}

	if length := len(skipPaths); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPaths {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request:      c.Request,
				DisableColor: conf.DisableColor,
				Keys:         c.Keys,
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			fmt.Fprint(out, formatter(param))
		}
	}
}
