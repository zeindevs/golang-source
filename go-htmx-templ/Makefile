run: build
	@./bin/app

build:
	@go build -o bin/app .

test:
	@go test -v ./...

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch

templ:
	templ generate --watch --proxy=http://localhost:3000
