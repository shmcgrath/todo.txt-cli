.DEFAULT_GOAL := build

.PHONY: format
format:
	go fmt ./...

.PHONY: vet
vet: format
	go vet ./...

.PHONY: build
build: vet
	go build ./cmd/todo
