{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"encoding/json"
	"time"
)

type WebsocketRealTimeMessage struct {
	Type    int    `json:"type" example:"1"`
	Message string `json:"message" example:"以 | 分割字段"`
}

// RealTimeMessageBroadcast 实时推送
var RealTimeMessageBroadcast chan []byte

func init() {
	RealTimeMessageBroadcast = make(chan []byte)

	go func() {
		message, _ := json.Marshal(WebsocketRealTimeMessage{
			Type:    1,
			Message: "golang|python|cpp",
		})

		for {
			RealTimeMessageBroadcast <- message
			time.Sleep(3 * time.Second)
		}
	}()
}
