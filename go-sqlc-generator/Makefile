all: build run

run:
	./bin/main

build:
	go build -o bin/main main.go

sqlc-generate:
	sqlc generate

dbinit:
	psql -U postgres -d sqlcgenerator < schema.sql
