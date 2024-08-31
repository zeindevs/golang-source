run: build
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main main.go
	
test:
	go test ./...

testr:
	go test ./... --race
	
testv:
	go test -v ./...

testvr:
	go test -v ./... --race
