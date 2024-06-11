.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z%\\\/_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

BIN_DIR := bin
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/$(BIN_DIR))

CONTROLLER_GEN_BIN := controller-gen
CONTROLLER_GEN_VER := v0.15.0
CONTROLLER_GEN_PKG := sigs.k8s.io/controller-tools/cmd/controller-gen

MANIFESTS_DIR := manifests

$(CONTROLLER_GEN_BIN): # Build controller-gen from tools folder.
	go install "${CONTROLLER_GEN_PKG}@${CONTROLLER_GEN_VER}"

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN_BIN) # Generates all manifests.
	@mkdir -p ${MANIFESTS_DIR}
	$(MAKE) clean-manifests
	$(CONTROLLER_GEN_BIN) \
		paths=./api/... \
		crd:crdVersions=v1 \
		output:crd:dir=./config/crd/bases
	@kustomize build config/crd -o ${MANIFESTS_DIR}
	@rm manifests/v1_configmap_release-name-validation.yaml

.PHONY: clean-manifests
clean-manifests: # Deletes all generated manifests.
	@rm ./config/crd/bases/*.yaml
	@rm ./manifests/*.yaml

.PHONY: ginkgo
ginkgo: ## Install required ginkgo version
	@go get github.com/onsi/ginkgo/v2
	@go install github.com/onsi/ginkgo/v2/ginkgo

.PHONY: test
test: ginkgo ## Run unit tests
	@echo "====> $@"
	@ginkgo
