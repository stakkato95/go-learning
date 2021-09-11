.DEFAULT_GOAL := test

fmt:
	goimports -l -w .
.PHONY:fmt

tets: fmt
	go test ./...
.PHONY:test