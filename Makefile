TEST_OPT=""

.PHONY: bootstrap
bootstrap:
	mkdir -p bin
	GOBIN=$(PWD)/bin go install github.com/golang/mock/mockgen@latest

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
lint:
	golangci-lint run --config ".github/.golangci.yml" --fast

.PHONY: build
build:
	go build -o ./bin/volcago ./cmd/volcago
