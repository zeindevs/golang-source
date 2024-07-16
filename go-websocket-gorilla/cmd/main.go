package main

import (
	"context"
	"fmt"
	"go-websocket-gorilla/internal"
	"log"
	"net/http"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	setupAPI(ctx)

	log.Println("Server up and listening on http://localhost:8080")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func setupAPI(ctx context.Context) {

	manager := internal.NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/login", manager.LoginHandler)
	http.HandleFunc("/ws", manager.ServeWS)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, manager.CurrentClients())
	})
}
