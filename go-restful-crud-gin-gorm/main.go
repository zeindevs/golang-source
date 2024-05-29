package main

import (
	"net/http"
	"restful-crud-gin-gorm/config"
	"restful-crud-gin-gorm/controller"
	"restful-crud-gin-gorm/helper"
	"restful-crud-gin-gorm/model"
	"restful-crud-gin-gorm/repository"
	"restful-crud-gin-gorm/router"
	"restful-crud-gin-gorm/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	// database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	// service
	tagsService := service.NewTagsServiceImpl(tagRepository, validate)

	// controller
	tagsController := controller.NewTagsController(tagsService)

	// routes
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	log.Info().Msg("Server up and listening on port: 8888")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
