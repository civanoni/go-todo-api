package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var db *gorm.DB

func ConnectDB() *gorm.DB {
	once.Do(func() {
		_ = godotenv.Load()
		dsn := os.Getenv("DATABASE_URL")

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
	})

	return db
}
