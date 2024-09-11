package main

import (
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))
	slog.Info("server up and listening on", "addr", "http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
