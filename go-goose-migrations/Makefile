include .env

all: build run

run:
	./bin/main

build:
	go build -o bin/main .

goose-create:
	goose -dir db/migrations create $(filter-out $@,$(MAKECMDGOALS)) -s sql

goose-up:
	goose postgres $(DATABASE_URL) -dir db/migrations up

goose-down:
	goose postgres $(DATABASE_URL) -dir db/migrations down
