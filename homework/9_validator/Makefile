.DEFAULT_GOAL := test

fmt:
	goimports -l -w .
.PHONY:fmt

test: fmt
	go test ./...
.PHONY:test


testprofile:
	go test -coverprofile=cover.out
.PHONY:testprofile

coverager: testprofile
	go tool cover -html=cover.out
.PHONY:coverager


clear:
	rm -f cover.out
.PHONY:clear