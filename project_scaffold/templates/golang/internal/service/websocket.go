{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"encoding/json"
	"sync"
	"time"

    "{{GOLANG_MODULE}}/internal/query"
)

var (
	HubServer                    *Hub
	HubClient                    *Hub
	HubRealTimeMessageSubscriber *Hub
	MapClients                   sync.Map
)

const (
	Activation     int = iota + 1
	Authentication
)

const (
	KeepAlive     = "KeepAlive"
	GetPrivateKey = "GetPrivateKey"
	PostPublicKey = "PostPublicKey"
	GetPosition   = "GetPosition"
)

type WebsocketMessage struct {
	Cmd string `json:"cmd" example:"GetPrivateKey"`
	query.Response
}

type WebsocketRealTimeMessage struct {
	Type    int    `json:"type" example:"1"`
	Message string `json:"message" example:"以 | 分割字段"`
}

// RealTimeMessageBroadcast 实时推送
var RealTimeMessageBroadcast chan []byte

func init() {
	go func() {
		message, _ := json.Marshal(WebsocketRealTimeMessage{
			Type:    1,
			Message: "golang|python|cpp",
		})

		for {
			if HubRealTimeMessageSubscriber != nil {
				HubRealTimeMessageSubscriber.Broadcast <- message
			}

			time.Sleep(3 * time.Second)
		}
	}()
}
