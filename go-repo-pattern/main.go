package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/zeindevs/go-repo-pattern/mailer"
	"github.com/zeindevs/go-repo-pattern/storage"
)

type application struct {
	userService *UserService
}

func main() {
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Mailer
	mailtrapMailer, err := mailer.NewMailTrapClient(
		os.Getenv("MAILTRAP_API_KEY"),
		os.Getenv("MAILTRAP_FROM_EMAIL"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// storage
	storage := storage.NewSqlStorage(db)

	// Service
	userService := NewUserService(storage, mailtrapMailer)

	app := &application{
		userService: userService,
	}

	http.HandleFunc("POST /user", app.createUserHandler)

	log.Println("Server starting on :9090")

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func (a *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode Request
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create user
	if err := a.userService.Create(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
