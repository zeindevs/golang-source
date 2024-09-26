run: build
	./bin/gowatch

build:
	go build -ldflags="-s -w" -o bin/gowatch main.go

tailwind:
	tailwindcss -i style.css -o public/style.css -c tailwind.config.js -m
