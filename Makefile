CMD_DIR = ./cmd/hexlet-path-size
BINARY_NAME = hexlet-path-size

.PHONY: lint test build run full-flow

.DEFAULT_GOAL := build

lint:
	golangci-lint run

test:
	go test -v ./tests

build:
	mkdir -p bin
	go build -o bin/${BINARY_NAME} ${CMD_DIR}

full-flow: lint test build

run:
	go run ${CMD_DIR}/main.go