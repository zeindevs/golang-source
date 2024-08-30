package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var Bun *bun.DB

func CreateDatabase() (*sql.DB, error) {
	godotenv.Load()
	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("user=%s dbname=%s passwod=%s host=%s port=5432", dbuser, dbname, &dbpassword, &dbhost)
	)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Init() error {
	db, err := CreateDatabase()
	if err != nil {
		return err
	}

	Bun = bun.NewDB(db, pgdialect.New())
	Bun.AddQueryHook(bundebug.NewQueryHook())

	return nil
}
