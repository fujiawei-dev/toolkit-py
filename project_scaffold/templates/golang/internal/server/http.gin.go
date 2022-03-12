{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"

	"{{GOLANG_MODULE}}/internal/api"
	"{{GOLANG_MODULE}}/internal/middleware"
)

// Start the REST API server using the configuration provided.
func Start(ctx context.Context, ch chan error) {
	app := newApp()

	// Register HTTP route handlers.
	api.RegisterRoutes(app)

	// Create new HTTP server instance.
	server := &http.Server{
		Addr:         conf.InternalHttpHostPort(),
		Handler:      app,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	go func() {
		log.Info().Msgf("http: web server started on %s", conf.ExternalHttpHostPort())
		ch <- server.ListenAndServe()
	}()

	// Graceful HTTP server shutdown.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := server.Close(); err != nil {
		log.Error().Msgf("http: web server shutdown failed: %v", err)
	}
}

func newApp() (app *gin.Engine) {
	// Set HTTP server mode.
	gin.SetMode(conf.HttpMode())

	// Support coloring in Windows
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.DefaultErrorWriter = colorable.NewColorableStderr()

	// Create new HTTP router engine without standard middleware.
	app = gin.New()

	// Register custom middleware.
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			DisableColor: conf.DetachServer(),
			Output:       conf.LogWriter(),
		}))

	app.Use(middleware.Recovery(api.ErrorUnexpected))

	app.Use(middleware.Cors())

	return
}
