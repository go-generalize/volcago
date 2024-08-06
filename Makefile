TEST_OPT=""
GOLANGCI_LINT_VERSION := 1.50.1

.PHONY: bootstrap
bootstrap:
	mkdir -p bin
	GOBIN=$(PWD)/bin go install go.uber.org/mock/mockgen@latest

.PHONY: bootstrap_golangci_lint
bootstrap_golangci_lint:
	mkdir -p ./bin
	GOBIN=${PWD}/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION)

.PHONY: test
test: goimports
	go test ./... -v ${TEST_OPT}

.PHONY: goimports
goimports:
	cd /tmp && go install golang.org/x/tools/cmd/goimports@latest

.PHONY: code_clean
code_clean:
	cd generator/testfiles && rm -rf */*_gen.go

.PHONY: lint
lint: lint_golangci_lint

.PHONY: lint_golangci_lint
lint_golangci_lint:
	./bin/golangci-lint run --config=".github/.golangci.yml" --fast ./...

.PHONY: build
build:
	go build -o ./bin/volcago ./cmd/volcago

.PHONY: gen_samples
gen_samples: build bootstrap
	go generate ./examples
	go test ./generator
