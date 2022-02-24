{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"os"
	"syscall"

	"github.com/mattn/go-colorable"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/event"
	"{{GOLANG_MODULE}}/pkg/fs"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

var rootCommand = &cobra.Command{
	Use:              conf.AppName(),
	Short:            "This is a Web-Application",
	Version:          conf.Version(),
	TraverseChildren: true,
}

func NewAppCommand() (*cobra.Command, error) {
	conf.SetFlags(rootCommand.PersistentFlags())

	rootCommand.AddCommand(configCommand)
	rootCommand.AddCommand(startCommand)
	rootCommand.AddCommand(stopCommand)

	rootCommand.SetOut(colorable.NewColorableStdout())
	rootCommand.SetErr(colorable.NewColorableStderr())

	return rootCommand, nil
}

// childAlreadyRunning tests if a .pid file at filePath is a running process.
// it returns the pid value and the running status (true or false).
func childAlreadyRunning(filePath string) (pid int, running bool) {
	if !fs.Exists(filePath) {
		return pid, false
	}

	pid, err := daemon.ReadPidFile(filePath)
	if err != nil {
		return pid, false
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return pid, false
	}

	return pid, process.Signal(syscall.Signal(0)) == nil
}
