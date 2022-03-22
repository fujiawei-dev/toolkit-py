{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"

	"github.com/urfave/cli"

	"{{GOLANG_MODULE}}/internal/entity"
)

var ResetCommand = cli.Command{
	Name:  "reset",
	Usage: "Resets the database, clears the cache, and removes log & backup files",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "database, d",
			Usage: "reset database only",
		},
	},
	Action: resetAction,
}

// resetAction resets the index and removes sidecar files after confirmation.
func resetAction(ctx *cli.Context) error {
	if err := conf.Init(ctx); err != nil {
		return fmt.Errorf("config init failed: %v", err)
	}

	onlyResetDatabase := ctx.Bool("database")

	// Reset database?
	if onlyResetDatabase {
		return entity.ResetDatabase()
	}

	return entity.ResetDatabase()
}
