{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"fmt"
	"os"
	"syscall"

	"github.com/labstack/gommon/color"
	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli"

	"{{GOLANG_MODULE}}/pkg/fs"
)

// StopCommand registers the stop cli command.
var StopCommand = cli.Command{
	Name:    "stop",
	Aliases: []string{"down"},
	Usage:   "Stops the web server in daemon mode",
	Action:  stopAction,
}

// stopAction stops the daemon if it is running.
func stopAction(ctx *cli.Context) error {
	if err := conf.InitSettings(ctx); err != nil {
		return fmt.Errorf("config init failed: %v", err)
	}

	fmt.Printf("looking for pid in %s\n", conf.PidFile())

	if !fs.IsFile(conf.PidFile()) {
		return fmt.Errorf("%s does not exist or is not a file", conf.PidFile())
	}

	dc := daemon.Context{PidFileName: conf.PidFile()}

	child, err := dc.Search()

	if err != nil {
		return err
	}

	err = child.Signal(syscall.SIGTERM)

	if err != nil && err != os.ErrProcessDone {
		return err
	}

	ps, err := child.Wait()

	if err != nil {
		_ = fs.DeleteFile(conf.PidFile())

		fmt.Println("daemon exited successfully")

		if conf.DetachServer() {
			color.Printf("⇨ https server stopped on %s\n", color.Green(conf.ExternalHttpHostPort()))
		}

		return err
	}

	fmt.Printf("daemon[%v] exited[%v]? successfully[%v]?\n", ps.Pid(), ps.Exited(), ps.Success())

	return nil
}
