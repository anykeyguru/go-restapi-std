.PHONY: build
build:
	go build -v -o apiserver.exe ./cmd/apiserver


.DEFAULT_GOAL := build