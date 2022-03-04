{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"os"
	"syscall"

	"github.com/labstack/gommon/color"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"

	"{{GOLANG_MODULE}}/pkg/fs"
)

// stopCommand registers the stop cli command
var stopCommand = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down"},
	Short:   "Stop the web server in daemon mode",
	Run:     stopAction,
}

func stopAction(cmd *cobra.Command, args []string) {
	if err := conf.InitSettings(); err != nil {
		cmd.Printf("config init failed: %v\n", err)
		return
	}

	cmd.Printf("looking for pid in %s\n", conf.PidFile())

	if !fs.IsFile(conf.PidFile()) {
		cmd.Printf("%s does not exist or is not a file\n", conf.PidFile())
		return
	}

	dc := daemon.Context{PidFileName: conf.PidFile()}

	child, err := dc.Search()

	if err != nil {
		cmd.Println(err)
		return
	}

	err = child.Signal(syscall.SIGTERM)

	if err != nil && err != os.ErrProcessDone {
		cmd.Println(err)
		return
	}

	ps, err := child.Wait()

	if err != nil {
		_ = fs.DeleteFile(conf.PidFile())

		cmd.Println("daemon exited successfully")

		if conf.DetachServer() {
			color.Printf("⇨ https server stopped on %s\n", color.Green(conf.ExternalHttpHostPort()))
		}

		return
	}

	cmd.Println("daemon[%v] exited[%v]? successfully[%v]?\n", ps.Pid(), ps.Exited(), ps.Success())
}
