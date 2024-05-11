package main

import (
	"log"
	"net/http"

	"ginjwt/config"
	"ginjwt/controller"
	"ginjwt/helper"
	"ginjwt/model"
	"ginjwt/repository"
	"ginjwt/router"
	"ginjwt/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Inite Repository
	usersRepository := repository.NewUsersRepositoryImpl(db)

	// Init service
	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)

	// Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUserController(usersRepository)

	routes := router.NewRouter(&router.RouterConfig{
		UserRepository:           usersRepository,
		AuthenticationController: authenticationController,
		UserController:           usersController,
	})

	server := &http.Server{
		Addr:    ":" + loadConfig.ServerPort,
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
