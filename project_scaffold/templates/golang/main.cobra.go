{{GOLANG_HEADER}}

package main

import (
	"fmt"
	"os"

	"{{GOLANG_MODULE}}/internal/command"
)

// @title Swagger Example
// @version 1.0.0
// @description Automatically generate RESTful API documentation with Swagger 2.0 for Go.

// @schemes http
// @host localhost:8080
// @BasePath /api/v1
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
