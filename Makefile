GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/cover \

.PHONY: setup
setup: ## Install all the build and lint dependencies
	@echo "--> Installing tools"
	@go get -u $(GOTOOLS)

.PHONY: dep
dep: ## Install all import dependencies
	@echo "--> Installing dependencies"
	@dep ensure

.PHONY: test
test: ## Run all the tests
	@echo "--> Running tests"
	@echo 'mode: atomic' > coverage.txt && go list ./... | grep -v /vendor/ | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s {}'

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: ## gofmt and goimports all go files
	@echo "--> Running gofmt/goimports"
	@find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: lint
lint: ## Run all the linters
	@echo "--> Running linters"
	@golangci-lint run

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: generate
generate: ## Run go generate
	@echo "--> Running go generate"
	@go generate

.PHONY: build
build: ## Build
	@echo "--> Building ..."
	@go build -o bin/optional ./cmd/optional/main.go

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
