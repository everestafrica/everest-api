package scheduler

import (
	"context"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/external/asset"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"sync"

	"time"
)

type scheduler struct {
	news            services.INewsService
	stock           services.IStockService
	alert           services.ISettingsService
	acctTransaction services.IAccountTransactionService
	budget          services.IBudgetService
	crypto          services.ICryptoService
	subscription    services.ISubscriptionService
	db              *gorm.DB
}

type IScheduler interface {
	RegisterSchedulers()
}

func RegisterSchedulers() {
	s := scheduler{
		news:            services.NewNewsService(),
		stock:           services.NewStockService(),
		alert:           services.NewSettingsService(),
		acctTransaction: services.NewAccountTransactionService(),
		budget:          services.NewBudgetService(),
		crypto:          services.NewCryptoService(),
		subscription:    services.NewSubscriptionService(),
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
		var c []models.CryptoDetail
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
		var users []models.User
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
		var users []models.User
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
		var users []models.User
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
		var users []models.User
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

	// Define the maximum number of workers to run concurrently
	maxWorkers := 10

	// Create a channel to send and receive work requests
	workQueue := make(chan func(context.Context), maxWorkers)

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Create a context to cancel the workers if necessary
	ctx, cancel := context.WithCancel(context.Background())

	// Start the workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, workQueue)
	}

	sch.StartAsync()

	// Wait for the scheduler to stop
	defer sch.Stop()

	// Wait for all workers to finish
	wg.Wait()

	// Cancel the context to stop any running workers
	cancel()
}

func GetTwoDaysLater() time.Time {
	return time.Now().AddDate(0, 0, 2)
}
func GetTomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// worker is a function that processes work requests from a channel
func worker(ctx context.Context, wg *sync.WaitGroup, workQueue chan func(context.Context)) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// The context has been cancelled, so exit the worker
			return
		case jobFunction, ok := <-workQueue:
			if !ok {
				// The work queue has been closed, so exit the worker
				return
			}

			// Perform the work for the cron job
			fmt.Printf("Starting cron job\n")
			jobFunction(ctx)
			fmt.Printf("Finished cron job\n")
		}
	}
}
