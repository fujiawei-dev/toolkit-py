{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"net/http"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"{{GOLANG_MODULE}}/internal/api"
	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/middleware"
)

// Start the REST API server using the configuration provided.
func Start(ctx context.Context, ch chan error) {
	app := newIrisApp()

	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,

		// Start the server and disable the default interrupt handler in order to
		// handle it clear and simple by our own, without any issues.
		DisableInterruptHandler: true,
	})

	registerRoutes(app)

	go func() {
		app.Logger().Infof("http: web server started on %s\n", conf.ExternalHttpHostPort())
		ch <- app.Run(iris.Addr(conf.InternalHttpHostPort()), config, iris.WithoutServerError(iris.ErrServerClosed))
	}()

	// Graceful HTTP server shutdown.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		app.Logger().Errorf("http: web server shutdown failed: %v", err)
	}
}

func registerRoutes(app *iris.Application) {
	app.OnErrorCode(http.StatusUnauthorized, api.AbortUnauthorized)
	app.OnErrorCode(http.StatusForbidden, api.AbortPermissionDenied)

	Websocket(app)

	router := app.Party(conf.BasePath())

	// Register the request id middleware
	router.UseRouter(requestid.New())

	// AllowAll create a new Cors handler with permissive configuration allowing all
	// origins with all standard methods with any header and credentials.
	router.UseRouter(cors.AllowAll())

	api.RegisterRoutes(app)
}

func newIrisApp() (app *iris.Application) {
	app = iris.New()

	app.Validator = form.Validator()

	app.Logger().SetLevel(conf.LogLevelString())
	app.Logger().SetOutput(conf.LogWriter())
	app.Logger().SetTimeFormat(conf.LogTimeFormat())

	app.UseRouter(middleware.Logger(app.Logger().Printer, conf.LogTimeFormat()))

	// Register the recovery, after accesslog and recover,
	// before end-developer's middleware.
	app.UseRouter(recover.New())

	return
}
