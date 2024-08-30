run: build 
	./bin/main

build:
	go build -o bin/main cmd/api/main.go

up:
	go run cmd/migrate/main.go up

down:
	go run cmd/migrate/main.go down

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

drop:
	go run cmd/drop/main.go

seed:
	go run cmd/seed/main.go
