all: build run

build:
	go build -o ./bin/main .

install:
	go install -ldflags="-s -w" .

run:
	./bin/main

clean:
	rm -rf ./bin