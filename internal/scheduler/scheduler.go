package scheduler

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"log"
	"time"
)

type scheduler struct {
	news   services.INewsService
	crypto services.ICryptoService
	db     *gorm.DB
	//account services.IAccountService
}

type IScheduler interface {
	RegisterSchedulers()
}

func RegisterSchedulers() {
	s := scheduler{
		news:   services.NewNewsService(),
		crypto: services.NewCryptoService(),
		db:     database.DB(),
	}

	sch := gocron.NewScheduler(time.UTC)

	//sch.Every(12).Hour().Do(func() {
	//	err := s.account.SetAccountTransactions()
	//})

	sch.Every(58).Minute().Do(func() {
		err := s.news.DeleteNews()
		if err != nil {
			log.Print(err)
		}
	})

	sch.Every(1).Hour().Do(func() {
		err := s.news.SetNews()
		if err != nil {
			log.Print(err)
		}
	})

	sch.Every(12).Hour().Do(func() {
		var c []models.CryptoDetail
		if err := s.db.Find(&c).Error; err != nil {
			log.Print(err)
		}
		for _, v := range c {
			err := s.crypto.SetWallet(v.Symbol, v.WalletAddress, v.UserId)
			if err != nil {
				log.Print(err)
				return
			}
		}
	})

	sch.StartAsync()
}
