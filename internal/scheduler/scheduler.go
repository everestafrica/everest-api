package scheduler

import (
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/model"
	"github.com/everestafrica/everest-api/internal/service"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"time"
)

type scheduler struct {
	news            service.INewsService
	stock           service.IStockService
	alert           service.ISettingsService
	acctTransaction service.IAccountTransactionService
	budget          service.IBudgetService
	crypto          service.ICryptoService
	subscription    service.ISubscriptionService
	db              *gorm.DB
}

type IScheduler interface {
	RegisterSchedulers()
}

func RegisterSchedulers() {
	s := scheduler{
		news:            service.NewNewsService(),
		stock:           service.NewStockService(),
		alert:           service.NewSettingsService(),
		acctTransaction: service.NewAccountTransactionService(),
		budget:          service.NewBudgetService(),
		crypto:          service.NewCryptoService(),
		subscription:    service.NewSubscriptionService(),
		db:              database.DB(),
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

	sch.Every(1).Hour().Do(func() {
		err := s.stock.SetStockData()
		if err != nil {
			log.Error("set stock error", err)
			return
		}
	})

	sch.Every(12).Hour().Do(func() {
		var c []model.CryptoDetail
		if err := s.db.Find(&c).Error; err != nil {
			log.Error("fetch crypto details error", err)
			return
		}
		for _, v := range c {
			err := s.crypto.UpdateWallet(v.Symbol, v.WalletAddress, v.UserId)
			if err != nil {
				log.Error("update crypto wallet error", err)
				return
			}
		}
	})

	sch.Every(12).Hour().Do(func() {
		var users []model.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error", err)
			return
		}
		for _, v := range users {
			err := s.acctTransaction.SetAccountTransactions(v.UserId)
			log.Error("update account transactions error", err)
		}
	})

	sch.Every(1).Minute().Do(func() {
		var users []model.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error in subscriptions cron", err)
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
					fmt.Println("the sub: ", sub.NextPayment, user.Email)
					fmt.Println(GetTwoDaysLater().Day())
					//go channels.SendMail(&channels.Email{
					//	Type:      channels.Subscription,
					//	Recipient: user.Email,
					//	Subject:   "REMINDER: Subscription Due",
					//	Body:      "The following subscription(s) payment is due in two days",
					//	Data:      sub,
					//})
				}
				if sub.NextPayment.Day() == GetTomorrow().Day() {
					fmt.Println("the sub: "+
						"", sub.NextPayment, user.Email)
					if err = s.db.Delete(&sub).Error; err != nil {
						log.Error("delete sub err: ", err)
					}
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

	sch.Every(1).Minute().Do(func() {
		var users []model.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error in alerts cron", err)
			return
		}
		for _, user := range users {
			alerts, err := s.alert.GetAllPriceAlerts(user.UserId)
			if err != nil {
				log.Error("get price alerts error", err)
			}
			for _, alert := range *alerts {
				price, assetErr := asset.GetAssetPrice(alert.Asset, alert.IsCrypto)
				if assetErr != nil {
					return
				}
				if *price >= alert.Target {
					// TODO
					//Push to queue
					//go channels.SendMail(&channels.Email{
					//	Type:      channels.Subscription,
					//	Recipient: user.Email,
					//	Subject:   "REMINDER: Subscription Due",
					//	Body:      "The following subscription(s) payment is due in two days",
					//	Data:      sub,
					//})
					// Also a Push Notification
				}
			}
		}
	})

	sch.Every(10).Minute().Do(func() {
		var users []model.User
		if err := s.db.Find(&users).Error; err != nil {
			log.Error("fetch users error", err)
			return
		}
		now := time.Now()
		year, month, _ := now.Date()

		firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
		for _, user := range users {
			outflow, err := s.acctTransaction.GetOutflow(types.DateRange{
				From: firstOfMonth,
				To:   now,
			}, user.UserId)
			if err != nil {
				log.Error("error from fetching outflow", err)
				return
			}
			monthBudget, err := s.budget.GetBudget(string(rune(month)), year, user.UserId)
			if err != nil {
				log.Error("error from budget", err)
				return
			}

			if int(outflow.Total) >= monthBudget.Amount {
				// Push Notification or Email
			}
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
