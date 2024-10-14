package main

import (
	"context"
	"log"
	"net/http"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "userId", 1)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("POST /upload", authenticate(http.HandlerFunc(storeUpload)))

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server listening on port :8080")
	srv.ListenAndServe()
}
