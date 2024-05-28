package database

import (
	"fmt"

	"go-crud-redis/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionPostgresDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Connected Successfully to the database (MySQL)")

	return db
}
