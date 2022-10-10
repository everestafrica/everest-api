package scheduler

import (
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/services"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"

	"time"
)

type scheduler struct {
	news    services.INewsService
	account services.IAccountTransactionService
	crypto  services.ICryptoService
	db      *gorm.DB
}

type IScheduler interface {
	RegisterSchedulers()
}

func RegisterSchedulers() {
	s := scheduler{
		news:    services.NewNewsService(),
		account: services.NewAccountTransactionService(),
		crypto:  services.NewCryptoService(),
		db:      database.DB(),
	}

	sch := gocron.NewScheduler(time.UTC)

	sch.Every(58).Minute().Do(func() {
		err := s.news.DeleteNews()
		if err != nil {
			log.Error("delete news error", err)
			return
		}
	})

	sch.Every(1).Hour().Do(func() {
		err := s.news.SetNews()
		if err != nil {
			log.Error("set news error", err)
			return
		}
	})

	sch.Every(12).Hour().Do(func() {
		var c []models.CryptoDetail
		if err := s.db.Find(&c).Error; err != nil {
			log.Error("fetch crypto details error", err)
			return
		}
		for _, v := range c {
			err := s.crypto.UpdateWallet(v.Symbol, v.WalletAddress, v.UserId)
			if err != nil {
				log.Error("update crypto details and txn error", err)
				return
			}
		}
	})

	sch.Every(12).Hour().Do(func() {
		var users []models.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error", err)
			return
		}
		for _, v := range users {
			err := s.account.SetAccountTransactions(v.UserId)
			log.Error("update account transactions error", err)
		}
	})

	sch.StartAsync()
}
