run: build
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main cmd/web/*.go

goose:
	goose -dir=assets/migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

migrate-up:
	goose -dir=assets/migrations sqlite3 app.db up

migrate-down:
	goose -dir=assets/migrations sqlite3 app.db down
