run: build
	@./bin/main.exe

build:
	@go build -o bin/main.exe
	
test:
	@go test ./...

testr:
	@go test ./... --race
	
testv:
	@go test -v ./...

testvr:
	@go test -v ./... --race