package main

import (
	"hex-structure/internal/adapters/primary/web"
	"hex-structure/internal/adapters/secondary/postgres"
	"hex-structure/internal/core/services/users"
	"log"
)

func main() {
  // Initialise secondary part implementations (Secodary adapters)
	userRepo, err := postgres.NewUserRepo() // <- this is swappable since its just a repo implementations
  if err != nil {
    log.Fatal("failed to init posgres user repo: %w", err)
  }

  // Initialise core service layer
  usersService := users.NewService(userRepo) // core business logic doesn't change.

  // Init primary (driving) adapter
  // this is swapple since we can spin up another primary adapter, and inject business logic
  srv := web.NewApp(usersService, web.WithPort(8000))

  go srv.Run()
}
