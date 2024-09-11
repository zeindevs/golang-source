package main

import (
	"encoding/json"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
				return
			}
			writeJSON(w, http.StatusInternalServerError, ErrInternalServer)
		}
	}
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return ErrMethodInvalid
	}

	user := User{
		ID:    1,
		Valid: true,
	}
	if !user.Valid {
		return ErrUserInvalid
	}

	return writeJSON(w, http.StatusOK, user)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
