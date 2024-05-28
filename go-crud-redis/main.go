package main

import (
	"fmt"
	"log"

	"go-crud-redis/config"
	"go-crud-redis/controller"
	"go-crud-redis/database"
	"go-crud-redis/model"
	"go-crud-redis/repo"
	"go-crud-redis/router"
	"go-crud-redis/usecase"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {

	// load config environment variables
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	// mysql
	db := database.ConnectionPostgresDB(&loadConfig)
	db.AutoMigrate(&model.Novel{})

	// redis
	rdb := database.ConnectionRedisDb(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	novelRepo := repo.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUseCase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)

	routes := router.NewRouter(app, novelController)

	fmt.Println("Server up and listening on port 5000")

	err := routes.Listen(":5000")
	if err != nil {
		panic(err)
	}
}
