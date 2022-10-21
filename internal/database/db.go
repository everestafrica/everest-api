package database

import (
	"github.com/everestafrica/everest-api/internal/models"
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

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		models.AccountDetail{},
		models.AccountTransaction{},
		models.Budget{},
		models.CryptoDetail{},
		models.CryptoTransaction{},
		models.Debt{},
		models.MonoUser{},
		models.News{},
		models.Subscription{},
		models.User{},
	)
}
