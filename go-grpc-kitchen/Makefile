run-orders:
	@./bin/orders
	# @go run services/orders/*.go

run-kitchen:
	@./bin/kitchen
	# @go run services/kitchen/*.go

build-orders:
	@go build -o bin/orders services/orders/*.go

build-kitchen:
	@go build -o bin/kitchen services/kitchen/*.go

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto"\
		--go_out=services/common/genproto/orders\
		--go_opt=paths=source_relative\
		--go-grpc_out=services/common/genproto/orders\
		--go-grpc_opt=paths=source_relative
