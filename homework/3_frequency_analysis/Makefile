.DEFAULT_GOAL := build

fmt:
	goimports -l -w .
.PHONY:fmt

test: fmt
	go test ./...
.PHONY:test

build: test
	go build run/main.go
.PHONY:build

run: build
	./main
.PHONY:run


coverprofile:
	go test -coverprofile=cover.out ./...
.PHONY:coverprofile

coverhtml: coverprofile
	go tool cover -html=cover.out
.PHONY:coverhtml


clean:
	rm main
.PHONY:clean

cleancover:
	rm cover.out
.PHONY:cleancover