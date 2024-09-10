all: build run

run: 
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main .

test:
	go test -v ./...

testr:
	go test -v ./... --race
