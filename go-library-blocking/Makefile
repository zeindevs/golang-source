run: build
	./bin/main

build:
	g++ -Wall -Wextra -o bin/main example/main.cpp -L./lib -lmain -lpthread -Wl,-rpath=./lib

build-release:
	g++ -Wall -Wextra -s -O3 -o bin/main example/main.cpp -L./lib -lmain -lpthread -Wl,-rpath=./lib

build-exe:
	x86_64-w64-mingw32-g++ -Wall -Wextra -o bin/main.exe example/main.cpp -L./lib -lmain -lpthread -Wl,-rpath=./lib

build-exe-release:
	x86_64-w64-mingw32-g++ -Wall -Wextra -s -O3 -o bin/main.exe example/main.cpp -L./lib -lmain -lpthread -Wl,-rpath=./lib

libso:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lib/libmain.so -buildmode=c-shared main.go

libdll:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o lib/libmain.dll -buildmode=c-shared main.go


