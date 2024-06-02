run: build
	@bin/main.exe

build:
	@go build -o bin/main.exe main.go

remote: build-remote
	@bin/remote.exe

build-remote:
	@go build -o bin/remote.exe remote/main.go
