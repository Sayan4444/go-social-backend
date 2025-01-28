package database

import (
	"log"
	"time"

	"github.com/Sayan4444/go-social-backend/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	DB_ADDRESS := utils.GetEnvAsString("DB_ADDRESS")

	d, err := gorm.Open(postgres.Open(DB_ADDRESS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB, err := d.DB()

	if err != nil {
		panic("db not connected")
	}

	DB_MAX_OPEN_CONNECTIONS := utils.GetEnvAsInt("DB_MAX_OPEN_CONNECTIONS")
	DB.SetMaxOpenConns(DB_MAX_OPEN_CONNECTIONS)

	DB_MAX_IDLE_CONNECTIONS := utils.GetEnvAsInt("DB_MAX_IDLE_CONNECTIONS")
	DB.SetMaxIdleConns(DB_MAX_IDLE_CONNECTIONS)

	DB_MAX_IDLE_TIME := utils.GetEnvAsInt("DB_MAX_IDLE_TIME")
	DB.SetConnMaxIdleTime(time.Duration(DB_MAX_IDLE_TIME) * time.Minute)

	DB_MAX_LIFETIME := utils.GetEnvAsInt("DB_MAX_LIFETIME")
	DB.SetConnMaxLifetime(time.Duration(DB_MAX_LIFETIME) * time.Minute)

	db = d
	log.Printf("DB Connected")
}

func GetDB() *gorm.DB {
	return db
}
