package main

import (
	"flag"
	"net/http"

	"github.com/zeindevs/go-folder-structure/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":49999", "todo")
	flag.Parse()

	http.HandleFunc("/user", api.HandleGetUser)
	http.HandleFunc("/account", api.HandleGetAccount)

	http.ListenAndServe(*listenAddr, nil)
}
