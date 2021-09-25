# suppress output, run `make XXX V=` to be verbose
V := @

default: test

.PHONY: test
test: GO_TEST_FLAGS := -race
test:
	$(V)go test -mod=vendor $(GO_TEST_FLAGS) --tags=$(GO_TEST_TAGS) ./...

.PHONY: lint
lint:
	$(V)golangci-lint run --config .golangci.yml
