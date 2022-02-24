{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/urfave/cli"
)

// VersionCommand registers the version cli command.
var VersionCommand = cli.Command{
	Name:   "config",
	Usage:  "Show settings from different sources",
	Action: configAction,
}

func configAction(ctx *cli.Context) error {
	conf.InitSettings(ctx)
	conf.Print()

	return nil
}
