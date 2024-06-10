run: build
	@./bin/main.exe

build:
	@go build -o bin/main.exe .

gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative msg/messages.proto
