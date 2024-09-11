run: build
	./bin/main

build:
	go build -ldflags="-s -w" -o bin/main main.go

test:
	go test -v ./... -count=1
