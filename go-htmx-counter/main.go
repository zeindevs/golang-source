package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	tmpStr := `<span id="counter">{{.Counter}}</span>`

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	counter := &Counter{}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		data := map[string]int{
			"Counter": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	r.Post("/increase", func(w http.ResponseWriter, r *http.Request) {
		counter.Increase()
		tmpl := template.Must(template.New("counter").Parse(tmpStr))
		data := map[string]int{
			"Counter": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})

	r.Post("/decrease", func(w http.ResponseWriter, r *http.Request) {
		counter.Decrease()
		tmpl := template.Must(template.New("counter").Parse(tmpStr))
		data := map[string]int{
			"Counter": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})

	log.Println("server listening on port :9000")
	http.ListenAndServe(":9000", r)
}
