all: build run

build:
	go build -o ./bin/main .

install:
	go install -ldflags="-s -w" .

release:
	go build -ldflags="-s -w" -o ./bin/main .

run:
	./bin/main

clean:
	rm -rf ./bin
