package main

import (
	"log"

	"github.com/zeindevs/go-realtime-chat/internal/user"
	"github.com/zeindevs/go-realtime-chat/internal/ws"
	"github.com/zeindevs/go-realtime-chat/router"
	"github.com/zeindevs/go-realtime-chat/server/db"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %v", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
  go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:9001")
}
