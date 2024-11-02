package main

import (
	"database/sql"
	"log"

  _ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://postgres:root@localhost:5432/goosemigrations?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Successfully connected to the database")

	rows, err := db.Query("SELECT id, name, city, coach_id FROM team")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, coach_id int
		var name, city string
		if err := rows.Scan(&id, &name, &city, &coach_id); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, City: %s, Coach: %d\n", id, name, city, coach_id)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}
}
