run: build
	@./bin/main.exe

build:
	@go build -o bin/main.exe

build-dev:
	@go build -tags dev -o bin/main.exe

build-prod:
	@go build -tags prod -o bin/main.exe

test:
	@go test -v ./...
