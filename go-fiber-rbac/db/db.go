package db

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *viper.Viper, log *logrus.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.GetString("DATABASE_URI")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Info("database connected")

	return db
}
