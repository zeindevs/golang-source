package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	listenAddr = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

var app *http.ServeMux

func init() {
	app = http.NewServeMux()
	app.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]any{"msg": "Hello Bro!"})
	})
	app.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]any{"msg": "Pong Bro!"})
	})
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
