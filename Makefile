gen:
	protoc --go_out=./grpc --go_opt=paths=source_relative \
	--go-grpc_out=./grpc --go-grpc_opt=paths=source_relative \
	proto/main.proto