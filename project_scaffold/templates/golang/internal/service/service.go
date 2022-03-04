{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"

	"{{GOLANG_MODULE}}/internal/config"
	"{{GOLANG_MODULE}}/internal/event"
)

var (
	log  = event.Logger()
	conf = config.Conf()
)

type Service struct {
	cron *gocron.Scheduler
}

func Start(ctx context.Context) {
	srv := NewService()

	// srv.RegisterCornEvents(cornHttpGetExamples)

	srv.Run()
	<-ctx.Done()
	srv.Stop()
}

func NewService() *Service {
	return &Service{cron: gocron.NewScheduler(time.Local)}
}

func (s *Service) RegisterCornEvents(events ...func(cron *gocron.Scheduler)) {
	for i := range events {
		events[i](s.cron)
	}
}

func (s *Service) Run() {
	s.cron.StartAsync()
}

func (s *Service) Stop() {
	s.cron.Stop()
}
