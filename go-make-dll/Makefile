build-dll:
	go build -ldflags="-s -w" -o mylib.dll -buildmode=c-shared mylib.go

build-c:
	clang -O3 -m64 -o main.exe main.c
