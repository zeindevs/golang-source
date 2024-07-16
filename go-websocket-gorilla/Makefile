run: build
	./bin/main.exe

build:
	go build -ldflags="-s -w" -o bin/main.exe cmd/main.go

ssl:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key
	openssl req -new -x509 -sha256 -key server.key -out server.crt -batch -days 3650
	
