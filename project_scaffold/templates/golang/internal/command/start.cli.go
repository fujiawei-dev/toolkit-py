{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"errors"
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
		return fmt.Errorf("config init failed: %v", err)
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		return errors.New("server port must be a number between 1 and 65535")
	}

	if !daemon.WasReborn() && conf.DetachServer() {
		color.Printf("⇨ https server started on %s\n", color.Green(conf.ExternalHttpHostPort()))

		if pid, ok := childAlreadyRunning(conf.PidFile()); ok {
			return fmt.Errorf("daemon already running with process id %v\n", pid)
		}

		dc := daemon.Context{PidFileName: conf.PidFile()}

		child, err := dc.Reborn()

		if err != nil {
			return err
		}

		if child != nil {
			if !fs.Overwrite(conf.PidFile(), []byte(strconv.Itoa(child.Pid))) {
				return fmt.Errorf("failed writing process id to %s", conf.PidFile())
			}

			return fmt.Errorf("daemon started with process id %v", child.Pid)
		}

		defer func() {
			if err = dc.Release(); err != nil {
				log.Printf("daemon release %v", err)
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

		log.Print("http: shutting down web server...")

		if serverError == http.ErrServerClosed || serverError == nil {
			log.Print("http: web server shutdown complete")
		} else {
			log.Printf("http: web server closed unexpect: %v", serverError)
		}

	case serverError = <-serverClosedSignal:
		cancel()

		if serverError != nil {
			return serverError
		}
	}

	return nil
}
