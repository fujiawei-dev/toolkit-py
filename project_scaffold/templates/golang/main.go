{{GOLANG_HEADER}}

package main

import (
	"fmt"
	"os"

    {%- if CLI_FRAMEWORK==".cli" %}

	"github.com/urfave/cli"
    {%- endif %}

	"{{GOLANG_MODULE}}/internal/command"
    {%- if CLI_FRAMEWORK==".cli" %}
	"{{GOLANG_MODULE}}/internal/config"
    {%- endif %}
)

// @Title        Swagger Example API
// @Version      1.0.0
// @Description  Automatically generate RESTful API documentation with Swagger 2.0 for Go.

// @Schemes   http
// @Host      127.0.0.1:8080
// @BasePath  /api/v1

// @SecurityDefinitions.Basic   BasicAuth
// @SecurityDefinitions.ApiKey  ApiKeyAuth
// @In                          header
// @Name                        Authorization
func main() {
    {%- if CLI_FRAMEWORK==".cobra" %}
	onError := func(err error) {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	cmd, err := command.NewAppCommand()

	if err != nil {
		onError(err)
	}

	if err = cmd.Execute(); err != nil {
		onError(err)
	}
    {%- elif CLI_FRAMEWORK==".cli" %}
	app := cli.NewApp()
	app.Version = config.Conf().Version()
	app.EnableBashCompletion = true
	app.Flags = config.CommandFlags

	app.Commands = []cli.Command{
		command.StartCommand,
		command.StopCommand,
		command.ResetCommand,
		command.VersionCommand,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
    {%- endif %}
}
