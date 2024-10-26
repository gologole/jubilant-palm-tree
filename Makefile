run:
	go run  cmd/main.go -race
generate-proto:
	protoc --go_out=./internal/transport/grpc --go-grpc_out=./internal/transport/grpc ./api/proto/file.proto

gocilint:

