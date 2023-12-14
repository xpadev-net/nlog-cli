gen:
	protoc --go_out=./src/pkg --go_opt=paths=source_relative \
	--go-grpc_out=./src/pkg --go-grpc_opt=paths=source_relative \
	proto/main.proto