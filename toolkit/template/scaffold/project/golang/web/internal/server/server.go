package server

import (
	"{{ main_module }}/internal/config"
	"{{ main_module }}/internal/event"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)
