proto:
	protoc pb/*.proto --go_out=. --go-grpc_out=.

server:
	go run cmd/main.go
