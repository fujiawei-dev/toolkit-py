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
		log.Fatalf("config init failed: %v", err)
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		log.Fatal("server port must be a number between 1 and 65535")
	}

	if !daemon.WasReborn() && conf.DetachServer() {
		color.Printf("⇨ https server started on %s\n", color.Green(conf.ExternalHttpHostPort()))

		if pid, ok := childAlreadyRunning(conf.PidFile()); ok {
			log.Infof("daemon already running with process id %v\n", pid)
			return
		}

		dc := daemon.Context{PidFileName: conf.PidFile()}

		child, err := dc.Reborn()

		if err != nil {
			log.Fatalf("failed reborn %v", err)
		}

		if child != nil {
			if !fs.Overwrite(conf.PidFile(), []byte(strconv.Itoa(child.Pid))) {
				log.Fatalf("failed writing process id to %s", conf.PidFile())
			}

			log.Infof("daemon started with process id %v", child.Pid)

			return
		}

		defer func() {
			if err = dc.Release(); err != nil {
				log.Errorf("daemon release %v", err)
			}
		}()
	}

	// pass this context down the chain
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan error)

	// start web server
	go server.Start(ctx, ch)

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

	<-quit

	log.Info("http: shutting down web server...")

	cancel()

	if err := <-ch; err == http.ErrServerClosed || err == nil {
		log.Info("http: web server shutdown complete")
	} else {
		log.Errorf("http: web server closed unexpect: %v", err)
	}
}
