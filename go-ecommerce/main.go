package main

import (
	"database/sql"
	"log"

	"github.com/zeindevs/go-ecommerce/cmd/api"
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

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Sucessfully connected!")
}
