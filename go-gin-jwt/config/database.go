package config

import (
	"fmt"
	"ginjwt/helper"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

  db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
  helper.ErrorPanic(err)

  fmt.Println("Connected successfully to the database")
  return db
}
