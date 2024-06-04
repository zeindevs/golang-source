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

proto-gen:
	@protoc \
	    --proto_path=proto "proto/types.proto"\
	    --go_out=proto\
	    --go_opt=paths=source_relative\
	    --go-grpc_out=proto\
	    --go-grpc_opt=paths=source_relative