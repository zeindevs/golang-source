all: build run

run:
	@./bin/main

build:
	@CGO_ENABLED=0 go build -ldflags='-s -w' -o bin/main main.go
