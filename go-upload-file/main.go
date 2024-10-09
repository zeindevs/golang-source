package main

import (
	"log"
	"net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("POST /upload", storeUpload)

  srv := http.Server{
    Addr: ":8080",
    Handler: mux,
  }

  log.Println("Server listening on port :8080")
  srv.ListenAndServe()
}
