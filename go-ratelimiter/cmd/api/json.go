package api

import "net/http"

func writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	return nil
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	return nil
}
