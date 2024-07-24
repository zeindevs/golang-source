package main

import (
	"fmt"
	"net/http"

	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/handler"
	"github.com/zeindevs/gospotify/internal"
)

func main() {
	cfg := config.NewConfig()
	player := internal.NewPlayerService(cfg)
	auth := internal.NewAuthService(cfg)
	handler := handler.NewHandler(cfg, auth, player)

	s := &http.ServeMux{}

	s.HandleFunc("GET /", handler.HandleIndex)
	s.HandleFunc("GET /login", handler.HandleLogin)
	s.HandleFunc("GET /login/client", handler.HandleClientLogin)
	s.HandleFunc("GET /callback", handler.HandleCallback)
	s.HandleFunc("GET /playing", handler.HandlePlaying)

	fmt.Println("Server up and listening on http://localhost:9001")
	if err := http.ListenAndServe(":9001", s); err != nil {
		panic(err)
	}
}
