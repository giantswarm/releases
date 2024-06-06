.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z%\\\/_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: ginkgo
ginkgo: ## Install required ginkgo version
	@go get github.com/onsi/ginkgo/v2
	@go install github.com/onsi/ginkgo/v2/ginkgo

.PHONY: test
test: ginkgo ## Run unit tests
	@echo "====> $@"
	@ginkgo
