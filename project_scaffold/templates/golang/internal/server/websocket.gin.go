{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"github.com/gin-gonic/gin"

    "{{GOLANG_MODULE}}/internal/api"
    "{{GOLANG_MODULE}}/internal/service"
)

func init() {
	api.AddRouteRegistrar(WebsocketServer)
}

func WebsocketServer(router *gin.RouterGroup) {
	hub := newHub("example")
	hub.broadcast = service.RealTimeMessageBroadcast
	go hub.run()

	router.GET("/ws", func(c *gin.Context) {
		conn, err := crossSiteUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Error().Msgf("ws: %v", err)
			return
		}

		client := NewClient(hub, conn, nil, nil)
		client.hub.register <- client

		// Allow collection of memory referenced by the caller by doing all work in new goroutines.
		go client.writePump()
		go client.readPump()
	})
}
