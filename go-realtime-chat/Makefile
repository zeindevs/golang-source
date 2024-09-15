all: build run

run: 
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main main.go

test:
	go test -v ./...

testr:
	go test -v ./... --race

postgres-init:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner-root gochat

dropdb:
	docker exec -it postgres15 dropdb gochat

migrate-init-user:
	migrate create -ext sql -dir db/migrations add_user_table

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5432/gochat?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5432/gochat?sslmode=disable" -verbose down
