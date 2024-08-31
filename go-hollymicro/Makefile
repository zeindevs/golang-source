run-scraper: scraper
	./bin/scraper

run-store: store
	./bin/store

scraper:
	go build -ldflags="-s -w" -o bin/scraper cmd/scraper/main.go

store:
	go build -ldflags="-s -w" -o bin/store cmd/store/main.go

proto:
	protoc \
		--proto_path=types "types/types.proto"\
		--go_out=types\
		--go_opt=paths=source_relative\
		--go-grpc_out=types\
		--go-grpc_opt=paths=source_relative

