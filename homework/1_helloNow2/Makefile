.DEFAULT_GOAL := run

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

clean:
	rm main