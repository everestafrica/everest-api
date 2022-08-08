package scheduler

import (
	"github.com/go-co-op/gocron"
	//"github.com/everestafrica/everet-api/internal/services"
	"time"
)

type scheduler struct {
}

func RegisterSchedulers() {
	//s := scheduler{}

	sch := gocron.NewScheduler(time.UTC)

	sch.Every(1).Hour().Do(func() {
		//s.rateService.UpdateRates()
	})

	sch.Every(10).Minute().Do(func() {
		//s.reversalService.ReverseAll()
	})

	sch.StartAsync()
}
