run: build
	@./bin/main

build:
	@go build -o bin/main
	
test:
	@go test ./...

testr:
	@go test ./... --race
	
testv:
	@go test -v ./...

testvr:
	@go test -v ./... --race