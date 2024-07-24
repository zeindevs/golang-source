package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url, err := h.Auth.Login(h.cfg.CLIENT_ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) HandleClientLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res, err := h.Auth.ClientLogin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"data": res})
}

func (h *Handler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var code = r.URL.Query().Get("code")
	var state = r.URL.Query().Get("state")

	if state == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": "state required"})
		return
	}

	res, err := h.Auth.Callback(code, state)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": "state required"})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"data": res})
}
