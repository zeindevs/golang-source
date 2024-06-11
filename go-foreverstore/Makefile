run: build
	@./bin/fs.exe

build:
	@go build -o bin/fs.exe
	
test:
	@go test ./...

testr:
	@go test ./... --race
	
testv:
	@go test -v ./...

testvr:
	@go test -v ./... --race