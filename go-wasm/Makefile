run: build-wasm build
	bin/server.exe

build:
	go build -ldflags="-s -w" -o bin/server.exe server.go

build-wasm:
	tinygo build -o html/wasm.wasm --no-debug -scheduler=none -target wasm ./wasm/wasm.go
