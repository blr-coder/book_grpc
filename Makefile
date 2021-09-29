.PHONY: build
.DEFAULT_GOAL := build

build:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./api/v1/*.proto
