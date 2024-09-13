package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type User struct {
	Name  string
	Email string
}

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(path.Join("www", "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	app.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(path.Join("www", "user-details.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, &User{
			Name:  "Anthony",
			Email: "anthony@example.com",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("server up and listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", app))
}
