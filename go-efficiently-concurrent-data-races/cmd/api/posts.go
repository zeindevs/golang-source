package main

import "net/http"

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {

	var payload UpdatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		return
	}

	// if err := Validate.Struct(payload); err != nil {
	//   return
	// }

	if payload.Content != nil {
	}

	if payload.Title != nil {

	}
}
