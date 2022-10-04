.PHONY: lint
lint:
	golangci-lint run


.PHONY: test
test:
	gotest ./...

