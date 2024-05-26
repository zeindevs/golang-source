run: build
	./bin/server.exe

build:
	go build -o bin/server.exe .

run-client: build-client
	./bin/client.exe

build-client:
	go build -o bin/client.exe ./cmd/client.go

proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto

docker-fetcher:
	docker run -p 3000:3000 -d fetcher:1
