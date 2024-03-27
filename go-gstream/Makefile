testp:
	@go run cmd/testpublish/main.go

testc:
	@go run cmd/testconsumer/main.go

run: build
	@./bin/gstream.exe

build:
	@go build -o bin/gstream.exe
