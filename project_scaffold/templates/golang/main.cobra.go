{{GOLANG_HEADER}}

package main

import (
	"fmt"
	"os"

	"{{GOLANG_MODULE}}/internal/command"
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
}
