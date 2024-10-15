package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zeindevs/go-posty/internal/models/sqlite"
)

type app struct {
	posts *sqlite.PostModel
}

func main() {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connected")

	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Server listening on port: 8080")
	srv.ListenAndServe()
}
