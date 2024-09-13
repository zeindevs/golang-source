package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Addr     string
	Port     string
	DBName   string
	SSLMode  string
}

func NewPostgresStorage(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
