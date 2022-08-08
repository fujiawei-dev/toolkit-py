package server

import (
	"github.com/gin-gonic/gin"

    "{{ main_module }}/internal/api"
    "{{ main_module }}/internal/service"
)

func init() {
	api.AddRouteRegistrar(WebsocketServer)
}

func WebsocketServer(router *gin.RouterGroup) {
	hubMaster := service.WebsocketScheduler.FindOrCreateHub(service.Masters, false)

	router.GET("/ws/master/:master_id", gin.BasicAuth(gin.Accounts{
		"ad58e54c8d4": "4e7213c403618c4f646ef7"}), func(c *gin.Context) {
		masterId := c.Param("master_id")
		if hubMaster.ExistsClient(masterId) {
			log.Warn().Msgf("master_id[%s] is already in use", masterId)
		}

		conn, err := service.CrossSiteUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Error().Msgf("ws: %v", err)
			return
		}

		client := service.NewClient(masterId, hubMaster, conn, nil, nil)
		client.Run() // Allow collection of memory referenced by the caller by doing all work in new goroutines.
		hubMaster.Register(client)

		service.WebsocketScheduler.FindOrCreateHub(masterId, true)
	})

	// ========================================================================

	hubRealTimeMessageSubscriber := service.WebsocketScheduler.FindOrCreateHub(service.RealTimeMessageSubscribers, false)

	router.GET("/ws/message", func(c *gin.Context) {
		conn, err := service.CrossSiteUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Error().Msgf("ws: %v", err)
			return
		}

		client := service.NewClient(c.ClientIP(), hubRealTimeMessageSubscriber, conn, nil, nil)
		client.Run() // Allow collection of memory referenced by the caller by doing all work in new goroutines.
		hubRealTimeMessageSubscriber.Register(client)
	})
}
