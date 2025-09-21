include .env

all: build run

build:
	go build -o bin/main main.go

run:
	./bin/main

db-up:
	goose -dir db/migrations postgres ${DATABASE_URI} up

db-down:
	goose -dir db/migrations postgres ${DATABASE_URI} down

db-reset:
	goose -dir db/migrations postgres ${DATABASE_URI} reset

db-seeder:
	go run cmd/seeder/main.go
