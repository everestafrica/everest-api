package main

import (
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/database/redis"
	"log"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dbConnection, err := database.ConnectDB(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}

	err = database.MigrateAll(dbConnection)

	if err != nil {
		log.Fatalf("migration error: %v", err)
	}

	defer func() {
		sqlDB, _ := dbConnection.DB()

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(10)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Minute * 30)

		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}()

	// Create new Redis
	redis.NewClient(cfg.RedisURL, "Everest")

	err = Main(cfg)

	if err != nil {
		log.Fatalf("server error: %v", err)
	}

}
