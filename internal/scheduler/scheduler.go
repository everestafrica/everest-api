package scheduler

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"

	"time"
)

type scheduler struct {
	news         services.INewsService
	account      services.IAccountTransactionService
	crypto       services.ICryptoService
	subscription services.ISubscriptionService
	db           *gorm.DB
}

type IScheduler interface {
	RegisterSchedulers()
}

func RegisterSchedulers() {
	s := scheduler{
		news:         services.NewNewsService(),
		account:      services.NewAccountTransactionService(),
		crypto:       services.NewCryptoService(),
		subscription: services.NewSubscriptionService(),
		db:           database.DB(),
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
	sch.Every(1).Minute().Do(func() {
		var users []models.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error", err)
			return
		}
		for _, user := range users {
			subs, err := s.subscription.GetAllSubscriptions(user.UserId)
			if err != nil {
				log.Error("get all subscriptions error", err)
				return
			}
			//var dueSubscriptions []models.Subscription
			for _, sub := range *subs {
				if sub.NextPayment.Day() == GetTwoDaysLater().Day() {
					fmt.Println("the sub", sub, user.Email)
					//go channels.SendMail(&channels.Email{
					//	Type:      channels.Subscription,
					//	Recipient: user.Email,
					//	Subject:   "REMINDER: Subscription Due",
					//	Body:      "The following subscription(s) payment is due in two days",
					//	Data:      sub,
					//})
				}
				if sub.NextPayment.Day() == GetTomorrow().Day() {
					fmt.Println("the sub", sub, user.Email)
					//go channels.SendMail(&channels.Email{
					//	Type:      channels.Subscription,
					//	Recipient: user.Email,
					//	Subject:   "REMINDER: Subscription Due",
					//	Body:      "The following subscription(s) payment is due tomorrow",
					//	Data:      sub,
					//})

				}

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

func GetTwoDaysLater() time.Time {
	return time.Now().AddDate(0, 0, 2)
}
func GetTomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}
