package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, data any) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return fmt.Errorf("invalid json data")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
}
