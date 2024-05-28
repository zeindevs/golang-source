package database

import (
	"fmt"

	"go-crud-redis/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMySQLDb(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	fmt.Println("Connected Successfully to the database (MySQL)")

	return db
}
