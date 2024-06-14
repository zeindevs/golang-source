run: build
	@./bin/main.exe

build:
	@go build -ldflags="-s -w" -o bin/main.exe

test:
	@go test -v ./...

ssh:
	@ssh-keygen -f "/home/streamer/.ssh/known_hosts" -R "[localhost]:2222"
	@cat big.txt | ssh localhost -p 2222

scp:
	@ssh-keygen -f "/home/streamer/.ssh/known_hosts" -R "[localhost]:2222"
	@scp -p 2222 -r testfolder localhost:aa@dd.com

bascp:
	@ssh-keygen -f "/home/streamer/.ssh/known_hosts" -R "[localhost]:2222"
	@scp -p 2222 main.go localhost:aaa.com

.PHONY: ssh
