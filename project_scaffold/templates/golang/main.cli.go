{{GOLANG_HEADER}}

package main

import (
	"os"

	"github.com/urfave/cli"

	"{{GOLANG_MODULE}}/internal/command"
	"{{GOLANG_MODULE}}/internal/config"
)

// @Title        Swagger Example API
// @Version      1.0.0
// @Description  Automatically generate RESTful API documentation with Swagger 2.0 for Go.

// @Schemes   http
// @Host      localhost:8080
// @BasePath  /api/v1

// @SecurityDefinitions.ApiKey  ApiKeyAuth
// @In                          header
// @Name                        Authorization
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
