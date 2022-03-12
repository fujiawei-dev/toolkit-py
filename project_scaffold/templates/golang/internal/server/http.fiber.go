{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"{{GOLANG_MODULE}}/internal/api"
)

// Start the REST API server using the configuration provided.
func Start(ctx context.Context, ch chan error) {
	app := newApp()

	registerRoutes(app)

	go func() {
		log.Info().Msgf("http: web server started on %s", conf.ExternalHttpHostPort())
		ch <- app.Listen(conf.InternalHttpHostPort())
	}()

	// Graceful HTTP server shutdown.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Error().Msgf("http: web server shutdown failed: %v", err)
	}
}

func registerRoutes(app *fiber.App) {
	router := app.Group(conf.BasePath())

	// Register the request id middleware
	app.Use(requestid.New())

	// AllowAll create a new Cors handler with permissive configuration allowing all
	// origins with all standard methods with any header and credentials.
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Authentication",
		AllowCredentials: true,
		ExposeHeaders:    "X-Request-ID",
	}))

	api.RegisterRoutes(app)
}

func newApp() (app *fiber.App) {
	app = fiber.New(fiber.Config{
		ServerHeader:          "Fast Http",
		ErrorHandler:          nil,
		DisableStartupMessage: true,
		AppName:               conf.AppName(),
		EnablePrintRoutes:     false,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${ip}:${port} | ${method} ${path} | ${status} | ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Shanghai",
		Output:     conf.LogWriter(),
	}))

	// Register the recovery, after accesslog and recover,
	// before end-developer's middleware.
	app.Use(recover.New())

	return
}
