.DEFAULT_GOAL := build

lint:
	golangci-lint run

run: #можно добавить через build:fmt предварительный линтинг
	go run cmd/hexlet-path-size/main.go

build: #можно добавить через build:fmt предварительный линтинг
	mkdir -p bin
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
