package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := ConnectDB(); err != nil {
		log.Fatal(err)
	}
	defer CloseDB()
	if err := SetupDB(); err != nil {
		log.Fatal(err)
	}
	if err := ParseTemplate(); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Get("/", HandleGetTasks)
	r.Post("/tasks", HandleCreateTask)
	r.Put("/tasks", HandleOrderTasks)
	r.Put("/tasks/{id}/toggle", HandleToggleTask)
	r.Get("/tasks/{id}/edit", HandleEditTask)
	r.Put("/tasks/{id}", HandleUpdateTask)
	r.Delete("/tasks/{id}", HandleDeleteTask)

	log.Println("server listening on port :3000")
	http.ListenAndServe(":3000", r)
}
