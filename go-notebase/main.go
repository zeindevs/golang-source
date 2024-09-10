package main

import "log"

func main() {
	cfg := PostgresConfig{
		DBUser:     "postgres",
		DBPassword: "root",
		DBName:     "gonotebase",
		DBPort:     "5432",
		DBSSLMode:  "disable",
	}

	storage := NewPostgresStorage(cfg)

	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer(":3000", db)
	apiServer.Run()
}
