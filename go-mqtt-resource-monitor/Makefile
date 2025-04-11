all: build run

run:
	@./bin/main

build:
	@go build -ldflags='-s -w' -o bin/main main.go

release:
	@CGO_ENABLED=0 go build -ldflags='-s -w' -o bin/main main.go
