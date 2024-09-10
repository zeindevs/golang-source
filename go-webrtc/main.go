package main

import (
	"log"
	"net/http"

	"github.com/zeindevs/go-webrtc/handlers"
	"github.com/zeindevs/go-webrtc/cmd/server"
	"github.com/zeindevs/go-webrtc/service/room"
)

var (
	LISTEN_PORT = ":5001"
)

func main() {
  router := &http.ServeMux{}

  svc := server.NewServer(&server.ServerConfig{
    Addr: LISTEN_PORT,
    Router: router,
  }) 

  router.Handle("/", http.FileServer(http.Dir("./www")))

  rooms := room.NewRoomMap()

  wsHandler := handlers.NewWSHandler(rooms)
  wsHandler.RegisterRoutes(router)	

	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
