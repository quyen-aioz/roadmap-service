API=api
GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint

build: lint
	@go build -o bin/${API} ./app/*.go

run: build
	@bin/${API}

lint:
	@${GOLANGCI_LINT} run

lint-fix:
	@${GOLANGCI_LINT} run --fix