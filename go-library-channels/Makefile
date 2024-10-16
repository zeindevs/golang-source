run: build
	./bin/main

build:
	g++ -Wall -Wextra -o bin/main example/main.cpp -I./lib -L./lib -lmain -lpthread -Wl,-rpath=./lib
	g++ -Wall -Wextra -static -s -O3 -o bin/main-static example/main.cpp -I./lib -L./lib -lmain -lpthread -Wl,-rpath=./lib

libso:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lib/libmain.so -buildmode=c-shared main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lib/libmain.a -buildmode=c-archive main.go

libdll:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o lib/libmain.dll -buildmode=c-shared main.go
