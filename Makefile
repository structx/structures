
deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...