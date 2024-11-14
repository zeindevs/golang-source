include .env

all: build run

run:
	./bin/main

build:
	go build -o bin/main main.go

river-migrate:
	river migrate-up --database-url $(DATABASE_URL)
