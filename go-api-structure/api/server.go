package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zeindevs/go-api-structure/storage"
	"github.com/zeindevs/go-api-structure/util"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /user", s.handleGetUserByID)
	http.HandleFunc("DELETE /user/{id}", s.handleDeleteUserByID)
	http.HandleFunc("fooa", s.handleFooA)
	http.HandleFunc("foob", s.handleFooB)
	http.HandleFunc("fooc", s.handleFooC)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"data": user})
}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	user := s.store.Get(id)

	num := util.Round2Dec(10.11214)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"data": user, "number": num})
}
