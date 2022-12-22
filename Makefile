server:
	go run cmd/main.go

contract:
	protoc --go_out=. --go-grpc_out=. proto/*.proto