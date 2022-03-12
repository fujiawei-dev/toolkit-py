{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"{{GOLANG_MODULE}}/internal/api"
	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/query"
)

// Start the REST API server using the configuration provided.
func Start(ctx context.Context, ch chan error) {
	app := newApp()

	api.RegisterRoutes(app)

	go func() {
		app.Logger.Infof("http: web server started on %s\n", conf.ExternalHttpHostPort())
		ch <- app.Start(conf.InternalHttpHostPort())
	}()

	// Graceful HTTP server shutdown.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := app.Close(); err != nil {
		app.Logger.Errorf("http: web server shutdown failed: %v", err)
	}
}

func newApp() (app *echo.Echo) {
	app = echo.New()

	app.HidePort = true
	app.HideBanner = true
	app.Logger = event.Logger()
	app.Debug = conf.Debug()

	app.Validator = form.NewValidatorForEcho()
	app.HTTPErrorHandler = httpErrorHandler

	app.Use(middleware.Logger(), middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut,
			http.MethodPatch, http.MethodPost, http.MethodDelete,
		},
	}))

	return
}

func httpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	httpError, ok := err.(*echo.HTTPError)

	if ok {
		if httpError.Internal != nil {
			if herr, ok := httpError.Internal.(*echo.HTTPError); ok {
				httpError = herr
			}
		}
	} else {
		httpError = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	code := httpError.Code
	message := httpError.Message
	response := query.Response{Code: code, Msg: message.(string)}

	if conf.Debug() {
		response.Err = httpError.Error()
	}

	if err = c.JSON(code, response); err != nil {
		c.Logger().Error(err)
	}
}
