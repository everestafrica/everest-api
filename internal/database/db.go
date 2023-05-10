package database

import (
	"github.com/everestafrica/everest-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var gormDB *gorm.DB

func ConnectDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Established database connection")

	setDB(db)

	return db, nil
}

func setDB(db *gorm.DB) {
	gormDB = db
}

func DB() *gorm.DB {
	return gormDB
}

//I hope this helps!

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		model.AccountDetail{},
		model.AccountTransaction{},
		model.Budget{},
		model.CryptoDetail{},
		model.CryptoTransaction{},
		model.Debt{},
		model.MonoUser{},
		model.News{},
		model.Asset{},
		model.Stock{},
		model.Tracker{},
		model.Subscription{},
		model.PriceAlert{},
		model.Category{},
		model.CoinRate{},
		model.User{},
	)
}
