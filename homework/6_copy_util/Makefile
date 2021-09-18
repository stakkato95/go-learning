.DEFAULT_GOAL := build

fmt:
	goimports -l -w .
.PHONY:fmt

test: fmt
	go test ./...
.PHONY:test

build: test
	go build cmd/copyutil/main.go
.phony:build

run: build
	./main
.PHONY:run


coverprofile:
	go test -coverprofile=cover.out ./...
.PHONY:coverprofile

coverage: coverprofile
	go tool cover -html=cover.out
.PHONY:coverage


clear:
	rm -f main
	rm -f cover.out
.PHONY:clear
	