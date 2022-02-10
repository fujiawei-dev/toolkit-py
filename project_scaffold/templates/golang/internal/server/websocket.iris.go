{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func Websocket(router *iris.Application) {
	ws := websocket.New(
		websocket.DefaultGobwasUpgrader,

		// https://github.com/kataras/iris/tree/master/_examples/websocket/native-messages
		websocket.Events{
			websocket.OnNativeMessage: func(conn *websocket.NSConn, message websocket.Message) error {
				log.Debugf("websocket: [recv] %s", message.Body)

				if !conn.Conn.IsClient() {
					return conn.Conn.Socket().WriteText([]byte("OK"), 0)
				}

				return nil
			},
		},
	)

	ws.OnConnect = func(c *websocket.Conn) error {
		log.Infof("websocket: [%s] connected", c.ID())
		// 实时输出日志
		router.Logger().AddOutput(newWsOutput(c))
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Infof("websocket: [%s] disconnected", c.ID())
	}

	router.Get("/access/log", websocket.Handler(ws)).
		Use(iris.FromStd(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			next(w, r)
		})) // Register middlewares (on handshake):
}

type wsOutput struct {
	conn *websocket.Conn
}

func (w wsOutput) Write(p []byte) (n int, err error) {
	if w.conn != nil && !w.conn.IsClosed() {
		err = w.conn.Socket().WriteText(p, 0)
	} else {
		w.conn = nil
	}

	return len(p), err
}

func newWsOutput(conn *websocket.Conn) wsOutput {
	return wsOutput{conn: conn}
}
