package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zeindevs/go-project-structure/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tables := []string{
		"accounts",
		"auction_histories",
		"auctions",
		"companies",
		"recurring_revenues",
		"schema_migrations",
		"sell_side_companies",
		"analyses",
		"user_details",
		"trading_limits",
		"buy_side_companies",
		"funding_requests",
		"recurring_requests",
		"strategies",
	}

	for _, table := range tables {
		query := fmt.Sprintf("drop table if exists %s cascade", table)
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
