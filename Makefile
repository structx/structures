
deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...

rpc:
	protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    protos/raft/raft_service.proto