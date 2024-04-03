package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /user", makeHandler(handleGetUsers))
	router.HandleFunc("GET /user/{id}", makeHandler(handleGetUserByID))

	http.ListenAndServe(":3000", router)
}

type User struct {
	ID uuid.UUID
}

type APIError struct {
	Status int
	Msg    string
}

func (e APIError) Error() string {
	return e.Msg
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "timmy",
		}
	}
	return writeJSON(w, http.StatusOK, User{ID: id})
}

func getUsers() ([]User, error) {
	return []User{}, nil
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := getUsers()
	if err != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "timmy",
		}
	}
	return writeJSON(w, http.StatusOK, users)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHandler(h apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if e, ok := err.(APIError); ok {
				slog.Error("API error", "err", err, "status", e.Status)
				writeJSON(w, e.Status, e)
			}
		}
	}
}
