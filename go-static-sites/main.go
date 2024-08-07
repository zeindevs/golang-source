package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
