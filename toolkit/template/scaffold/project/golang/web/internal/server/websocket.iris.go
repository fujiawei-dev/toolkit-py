package server

import (
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"

	"{{ main_module }}/internal/api"
	"{{ main_module }}/internal/service"
)

type WebsocketWriter struct {
	conn *websocket.Conn
}

func (w WebsocketWriter) Write(p []byte) (n int, err error) {
	if w.conn != nil && !w.conn.IsClosed() {
		err = w.conn.Socket().WriteText(p, 0)
	} else {
		w.conn = nil
	}

	return len(p), err
}

func NewWebsocketWriter(conn *websocket.Conn) WebsocketWriter {
	return WebsocketWriter{conn: conn}
}

func WebsocketLogWriter(router iris.Party) {
	ws := websocket.New(websocket.DefaultGobwasUpgrader, websocket.Events{})

	ws.OnConnect = func(c *websocket.Conn) error {
		router.Logger().AddOutput(NewWebsocketWriter(c))
		return nil
	}

	router.Get("/access/log", websocket.Handler(ws)).Use(iris.FromStd(func(
		w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(w, r)
	}))
}

func WebsocketClient(router iris.Party) {
	ws := websocket.New(websocket.DefaultGobwasUpgrader, websocket.Events{})

	ws.OnConnect = func(c *websocket.Conn) error {
		if c.ID() != "" {
			service.MapClientConnections.Delete(c.ID())
		}

		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		if c.ID() == "" {
			c.Close()
			return
		}

		service.MapClientConnections.Store(c.ID(), c)
	}

	router.Get("/client/{code:string}", websocket.Handler(ws, func(c iris.Context) string {
		code := c.Params().Get("code")
		if code == "" {
			api.ErrorInvalidParameters(c, errors.New("code(string) is required"))
			return ""
		}

		return code
	})).Use(iris.FromStd(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(w, r)
	}))
}

func Websocket(router iris.Party) {
	WebsocketLogWriter(router)
	WebsocketClient(router)
}
