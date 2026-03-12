.PHONY: test build lint help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests
	go test ./...

build: ## Build the tmon binary
	go build -o tmon ./cmd/tmon

lint: ## Run static analysis
	go vet ./...
