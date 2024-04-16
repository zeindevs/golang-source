run: build
	@./bin/main.exe --listenAddr :5001

build:
	@go build -o bin/main.exe

test:
	@go test -v ./...

test-server:
	@go test -run TestServerWithMultiClients -v

test-client1:
	@go test ./client -run TestNewClient1 -v

test-redis-client:
	@go test ./client -run TestNewRedisClient -v
