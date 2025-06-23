all: build run

run:
	@./bin/main

build:
	@go build -ldflags='-s -w' -o bin/main main.go

templ:
	templ generate

tailwind:
	@tailwindcss -i static/main.css -o static/style.css -m
