package main

import (
	"fmt"
	"net/http"
	"time"
)

func gorotineClosure() {

	names := []string{"Percy", "Gopher", "Santa"}

	for _, v := range names {
		// Create a goroutine for a closure
		go func() {
			fmt.Println(v)
		}()
	}

  time.Sleep(1 *time.Second)
}

func httpMux()  {

  mux := http.NewServeMux()

  mux.HandleFunc("GET /hello/", func(w http.ResponseWriter, r*http.Request) { 
    w.Write([]byte("Hello"))
  })

  mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r*http.Request) { 
    w.Write([]byte(fmt.Sprintf("Hello %s!", r.PathValue("name"))))
  })

  fmt.Println("http server listening on port :8080")
  if err := http.ListenAndServe(":8080", mux); err != nil {
    panic(err)
  }
}

func main() {
  // gorotineClosure()
  httpMux()
}
