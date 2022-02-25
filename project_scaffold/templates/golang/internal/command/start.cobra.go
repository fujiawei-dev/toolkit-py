{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/labstack/gommon/color"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"

	"{{GOLANG_MODULE}}/internal/server"
	"{{GOLANG_MODULE}}/internal/service"
	"{{GOLANG_MODULE}}/pkg/fs"
)

// startCommand registers the start cli command
var startCommand = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up"},
	Short:   "Start the web server",
	Run:     startAction,
}

func startAction(cmd *cobra.Command, args []string) {
	if err := conf.Init(); err != nil {
		cmd.Printf("config init failed: %v", err)
		return
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		cmd.Printf("server port must be a number between 1 and 65535")
		return
	}

	if !daemon.WasReborn() && conf.DetachServer() {
		color.Printf("⇨ https server started on %s\n", color.Green(conf.ExternalHttpHostPort()))

		if pid, ok := childAlreadyRunning(conf.PidFile()); ok {
			cmd.Printf("daemon already running with process id %v\n", pid)
			return
		}

		dc := daemon.Context{PidFileName: conf.PidFile()}

		child, err := dc.Reborn()

		if err != nil {
			cmd.Printf("failed reborn %v", err)
			return
		}

		if child != nil {
			if !fs.Overwrite(conf.PidFile(), []byte(strconv.Itoa(child.Pid))) {
				cmd.Printf("failed writing process id to %s", conf.PidFile())
				return
			}

			cmd.Printf("daemon started with process id %v", child.Pid)

			return
		}

		defer func() {
			if err = dc.Release(); err != nil {
				cmd.Printf("daemon release %v", err)
			}
		}()
	}

	// pass this context down the chain
	ctx, cancel := context.WithCancel(context.Background())

	// start web server
	serverClosedSignal := make(chan error)
	go server.Start(ctx, serverClosedSignal)

	// start service
	go service.Start(ctx)

	// set up proper shutdown of daemon and web server
	quit := make(chan os.Signal)

	signal.Notify(
		quit,
		// kill -SIGINT XXXX or Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// os.Kill  is equivalent with the syscall.Kill
		os.Kill,
		syscall.SIGKILL, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)

	var serverError error

	select {
	case <-quit:
		cancel()

		serverError = <-serverClosedSignal

		log.Print("http: shutting down web server...")

		if serverError == http.ErrServerClosed || serverError == nil {
			log.Print("http: web server shutdown complete")
		} else {
			log.Printf("http: web server closed unexpect: %v", serverError)
		}

	case serverError = <-serverClosedSignal:
		cancel()

		if serverError != nil {
			cmd.Print(serverError)
		}
	}
}
