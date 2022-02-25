{{GOLANG_HEADER}}

package main

import (
	"os"

	"github.com/urfave/cli"

	"{{GOLANG_MODULE}}/internal/command"
	"{{GOLANG_MODULE}}/internal/config"
)

// @title Swagger Example
// @version 1.0.0
// @description Automatically generate RESTful API documentation with Swagger 2.0 for Go.

// @schemes http
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := cli.NewApp()
	app.Version = config.Conf().Version()
	app.EnableBashCompletion = true
	app.Flags = config.CommandFlags

	app.Commands = []cli.Command{
		command.StartCommand,
		command.StopCommand,
		command.VersionCommand,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
