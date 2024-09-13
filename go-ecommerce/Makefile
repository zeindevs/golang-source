all: build run

run: 
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main main.go

test:
	go test -v ./...

testr:
	go test -v ./... --race

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down
