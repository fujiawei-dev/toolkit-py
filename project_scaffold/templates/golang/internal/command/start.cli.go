{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/labstack/gommon/color"
	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli"

	"{{GOLANG_MODULE}}/internal/server"
	"{{GOLANG_MODULE}}/internal/service"
	"{{GOLANG_MODULE}}/pkg/fs"
)

// StartCommand registers the start cli command
var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Start the web server",
	Action:  startAction,
}

func startAction(ctx *cli.Context) error {
	if err := conf.Init(ctx); err != nil {
		fmt.Printf("config init failed, %v\n", err)
		return nil
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		fmt.Println("server port must be a number between 1 and 65535")
		return nil
	}

	if !daemon.WasReborn() && conf.DetachServer() {
		color.Printf("⇨ https server started on %s\n", color.Green(conf.ExternalHttpHostPort()))

		if pid, ok := childAlreadyRunning(conf.PidFile()); ok {
			fmt.Printf("daemon already running with process id %v\n", pid)
			return nil
		}

		dc := daemon.Context{PidFileName: conf.PidFile()}

		child, err := dc.Reborn()

		if err != nil {
			fmt.Printf("daemon reborn failed, %v\n", err)
			return nil
		}

		if child != nil {
			if !fs.Overwrite(conf.PidFile(), []byte(strconv.Itoa(child.Pid))) {
				fmt.Printf("failed writing process id to %s\n", conf.PidFile())
				return nil
			}

			fmt.Printf("daemon started with process id %v\n", child.Pid)

			return nil
		}

		defer func() {
			if err = dc.Release(); err != nil {
				fmt.Printf("daemon release %v\n", err)
			}
		}()
	}

	// pass this context down the chain
	cctx, cancel := context.WithCancel(context.Background())

	// start web server
	serverClosedSignal := make(chan error)
	go server.Start(cctx, serverClosedSignal)

	// start service
	go service.Start(cctx)

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

		fmt.Println("http: shutting down web server...")

		if serverError == http.ErrServerClosed || serverError == nil {
			fmt.Println("http: web server shutdown complete")
		} else {
			fmt.Printf("http: web server closed unexpect, %v\n", serverError)
		}

	case serverError = <-serverClosedSignal:
		cancel()

		if serverError != nil {
			fmt.Printf("http: web server started failed, %v\n", serverError)
			return nil
		}
	}

	return nil
}
