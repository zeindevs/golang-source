build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

deploy_prod: build
	serverless deploy --stage prod --aws-profile maxaldtools

