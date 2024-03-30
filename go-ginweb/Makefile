run: build
	@./bin/main.exe

build:
	@go build -o bin/main.exe

buildop:
	@go build -ldflags="-s -w" -o bin/main-optimize.exe

test:
	@go test -v ./...
