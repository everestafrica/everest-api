package scheduler

import (
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/go-co-op/gocron"
	"time"
)

type scheduler struct {
	news services.INewsService
}

func RegisterSchedulers() {
	s := scheduler{
		news: services.NewNewsService(),
	}

	sch := gocron.NewScheduler(time.UTC)

	sch.Every(59).Minute().Do(func() {
		s.news.DeleteNews()
	})
	sch.Every(1).Hour().Do(func() {
		s.news.SetNews()
	})

	sch.StartAsync()
}
