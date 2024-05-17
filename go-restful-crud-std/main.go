package main

import (
	"net/http"
	"restful-crud-std/config"
	"restful-crud-std/controller"
	"restful-crud-std/helper"
	"restful-crud-std/repository"
	"restful-crud-std/router"
	"restful-crud-std/service"

	"github.com/rs/zerolog/log"
)

func main() {
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	log.Info().Msg("Server up and listening on port :8888")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
