package scheduler

import (
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/go-co-op/gocron"
	"log"
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

	sch.Every(58).Minute().Do(func() {
		err := s.news.DeleteNews()
		if err != nil {
			log.Print(err)
			return
		}
	})

	//sch.Every(1).Hour().Do(func() {
	//	err := s.news.SetNews()
	//	if err != nil {
	//		log.Print(err)
	//		return
	//	}
	//})

	sch.StartAsync()
}
