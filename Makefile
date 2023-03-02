.PHONY: build
build:
	go build -v ./cmd/api
	
.PHONY: test
test:
	go test -v -timeout 30s ./...

.DEFAULT_GOAL := build