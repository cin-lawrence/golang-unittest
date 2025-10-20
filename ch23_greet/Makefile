.PHONY: build
build:
	golangci-lint run
	go test ./...

.PHONY: protogen
protogen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    adapters/grpcserver/greet.proto

.PHONY: unit
unit:
	go test -short ./...
