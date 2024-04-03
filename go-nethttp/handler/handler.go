package handler

import (
	"log"
	"net/http"

	"github.com/zeindevs/nethttp/middleware"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.AuthUserID).(string)
	if !ok {
		log.Println("invalid user ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println("creating comment for user:", userID)

	w.Write([]byte("coment created for " + userID))
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}
