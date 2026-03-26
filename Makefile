API=api
GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint

build: 
	@go build -o bin/${API} ./app/*.go

run: 
	@bin/${API}

lint:
	@${GOLANGCI_LINT} run

lint-fix:
	@${GOLANGCI_LINT} run --fix