package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

var database = map[string]string{
	"user": "password",
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		password := r.FormValue("password")

		if pass, ok := database[user]; !ok || pass != password {
			err := http.StatusUnauthorized
			http.Error(w, "Invalid username or password", err)
			return
		}

		next(w, r)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("User %s Hit Endpoint", r.FormValue("user"))
		next(w, r)
	}
}

var middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authMiddleware,
	loggingMiddleware,
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, welcome to my website!")
}

func panic(w http.ResponseWriter, r *http.Request) {
	var tmp *int
	*tmp += 1
}

func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				msg := "Caught panic: %v, Stack trace: %s"
				log.Printf(msg, err, string(debug.Stack()))

				er := http.StatusInternalServerError
				http.Error(w, "Internal Server Error", er)
			}

		}()
		next(w, r)
	}
}

func main() {
	h := welcomeHandler
	for _, m := range middleware {
		h = m(h)
	}
	http.HandleFunc("/welcome", h)
	http.HandleFunc("/panic", recoveryMiddleware(panic))
	// http.HandleFunc("/panic", panic)

	log.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
