package service

import (
    "{{ main_module }}/internal/query"
)

var (
	WebsocketScheduler *Scheduler
)

func init() {
	WebsocketScheduler = NewScheduler()
	go WebsocketScheduler.Run()
	WebsocketScheduler.FindOrCreateHub(Default, true)
}

const (
	Masters                    = "masters"
	Default                    = "default"
	RealTimeMessageSubscribers = "realtime_message_subscribers"
)

type WebsocketMessage struct {
	Cmd string `json:"cmd" example:"GetSessionKey"`
	query.Response
}
