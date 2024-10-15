package main

import (
	"html/template"
	"net/http"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {
	posts, err := app.posts.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles("./assets/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, map[string]any{"Posts": posts})
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func (app *app) storePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.posts.Insert(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
