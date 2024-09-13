package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/zeindevs/go-ecommerce/config"
	"github.com/zeindevs/go-ecommerce/db"
)

func main() {
	db, err := db.NewPostgresStorage(db.Config{
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Addr:     config.Envs.DBAddress,
		Port:     config.Envs.DBPort,
		DBName:   config.Envs.DBName,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
