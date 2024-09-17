include .env

MIGRATIONS_PATH = cmd/migrate/migrations

all: build run

run: 
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main cmd/api/*.go

test:
	go test -v ./...

testr:
	go test -v ./... --race

migrate:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) up

migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATOR_ADDR) down
