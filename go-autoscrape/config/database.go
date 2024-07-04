package config

import (
	"fmt"
	"log"

	"github.com/zeindevs/goautoscrape/internal/model"
	"github.com/zeindevs/goautoscrape/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Get("DB_HOST"), cfg.Get("DB_PORT"), cfg.Get("DB_USER"), cfg.Get("DB_PASS"), cfg.Get("DB_NAME"), cfg.Get("DB_SSLMODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.ErrorPanic(err)

	db.AutoMigrate(&model.Otakudesu{})

	log.Println("Successfully connected to database")

	return db
}
