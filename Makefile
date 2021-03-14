GOTOOLS = \
	github.com/golangci/golangci-lint/cmd/golangci-lint \

SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

.PHONY: setup
setup: ## Install all the build and lint dependencies
	@echo "--> Installing tools"
	go install $(GOTOOLS)

.PHONY: test
test: ## Run all the tests
	@echo "--> Running tests"
	go test $(TEST_OPTIONS) -failfast $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=1m

.PHONY: fmt
fmt: ## gofmt and goimports all go files
	@echo "--> Running gofmt/goimports"
	@find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: lint
lint: ## Run all the linters
	@echo "--> Running linters"
	golangci-lint run

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: generate
generate: ## Run go generate
	@echo "--> Running go generate"
	go generate

.PHONY: build
build: ## Build
	@echo "--> Building ..."
	go build -o bin/optional ./cmd/optional/main.go

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
