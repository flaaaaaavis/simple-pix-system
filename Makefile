generate:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/user/v1/user.proto
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/pix/v1/pix.proto

build:
	@echo "........Building application........"
	@go build -o user-bin src/user/*.go
	@go build -o pix-bin src/pix/*.go

run:
	@echo "........Running User........"
	@go run src/user/*.go

run_pix:
	@echo "........Running Pix........"
	@go run src/pix/*.go