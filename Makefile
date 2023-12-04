generate:
	@protoc --proto_path=protobuf --go_out=plugind=grpc:protobuf/gen --go_opt=source_relative user/v1/user.proto
	@protoc --proto_path=protobuf --go_out=plugind=grpc:protobuf/gen --go_opt=source_relative pix/v1/pix.proto

build:
	@echo "........Building application........"
	@go build -o user-bin src/user/*.go
	@go build -o pix-bin src/pix/*.go

run:
	@echo "........Running User........"
	@go run src/user/*.go

run_pix:
	@echo "........Running User........"
	@go run src/pix/*.go