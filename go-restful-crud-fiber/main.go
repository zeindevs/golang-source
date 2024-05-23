package main

import (
	"log"
	"restful-crud-fiber/config"
	"restful-crud-fiber/controller"
	"restful-crud-fiber/model"
	"restful-crud-fiber/repository"
	"restful-crud-fiber/router"
	"restful-crud-fiber/service"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variable", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewServiceImpl(noteRepository, validate)

	// Init Controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Println("Server up and listening on port :8888")

	log.Fatal(app.Listen(":8888"))
}
