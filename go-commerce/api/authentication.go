package api

import (
	"encoding/json"
	"net/http"
)

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationHandler struct {
}

func (h *AuthenticationHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authReq := AuthenticationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]any{"message": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]any{"token": nil})
}
