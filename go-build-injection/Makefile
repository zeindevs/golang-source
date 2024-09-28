run: build
	./bin/main

build:
	go build -ldflags="-X main.buildDate=$$(date +%Y%m%d%H%M%S) -X main.commitHash=$$(git rev-parse main)" -o bin/main main.go
