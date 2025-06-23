package main

import (
	"log"
	"net/http"

	"github.com/zeindevs/go-htmx-websocket/db"
	"github.com/zeindevs/go-htmx-websocket/handler"
	"github.com/zeindevs/go-htmx-websocket/live"
	"github.com/zeindevs/go-htmx-websocket/router"
)

func main() {
	bundb := db.NewBunDB()

	notification := live.NewNotification()
	go notification.Run()

	handler := handler.NewHandler(bundb, notification)

	r := router.NewRouter(handler, http.Dir("./static"))

	log.Println("server listening on port http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
