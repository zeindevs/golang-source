all: build run

run:
	./bin/main

build:
	go build -o bin/main cmd/main.go

templ:
	templ generate

tailwind:
	tailwindcss -i static/custom.css -o static/css/style.css -c tailwind.config.js -m
