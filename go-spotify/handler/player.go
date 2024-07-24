package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/zeindevs/gospotify/types"
)

func (h *Handler) HandlePlaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	file, err := os.ReadFile("secret.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	var secret types.Secret
	json.Unmarshal(file, &secret)

	data, err := h.Player.GetCurrentPlaying(secret.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"data": data})
}
