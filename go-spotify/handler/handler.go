package handler

import (
	"encoding/json"
	"net/http"

	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/internal"
)

type Handler struct {
	cfg    *config.Config
	Auth   *internal.AuthService
	Player *internal.PlayerService
}

func NewHandler(cfg *config.Config, auth *internal.AuthService, player *internal.PlayerService) *Handler {
	return &Handler{
		cfg:    cfg,
		Auth:   auth,
		Player: player,
	}
}

func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"index":    "http://localhost:9001/",
		"login":    "http://localhost:9001/login",
		"client":   "http://localhost:9001/login/client",
		"callback": "http://localhost:9001/callback",
		"playing":  "http://localhost:9001/playing",
	}

	json.NewEncoder(w).Encode(map[string]any{"data": data})
}
