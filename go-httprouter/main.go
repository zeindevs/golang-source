package main

import (
	"log/slog"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /user/{id}", handleGetUserByID)
  
  slog.Info("Server started in", "port", 3000)
	http.ListenAndServe(":3000", router)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
  _ = id
	// slog.Info("receive request", "id", id)
}
