.PHONY: build
build:
	go build -v -o apiserver.exe ./cmd/apiserver

.PHONY: test
test:
	go test -v -timeout 30s ./...  
	

.DEFAULT_GOAL := build