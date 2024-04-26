package handlers

import (
	"net/http"

	"go-htmx-templ/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
