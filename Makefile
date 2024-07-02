PROJECT_NAME = regexp_utility
OUTPUT_PATH ?= bin/$(PROJECT_NAME)

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit: vet lint

## vet: examine Go source code and report suspicious constructs
.PHONY: vet
vet:
	@go mod verify
	@go vet ./...

## lint: run various Go source code linters
.PHONY: lint
lint: golangci-lint
	@$(GOLANGCI_LINT) run

## lint.fix: run various Go source code linters and automatically fix warnings
.PHONY: lint.fix
lint.fix: golangci-lint
	@$(GOLANGCI_LINT) run --fix

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	@go test -v -race -buildvcs ./...


## build: build the application
.PHONY: build
build:
	@mkdir -p $(dir $(OUTPUT_PATH))
	@go build -o $(OUTPUT_PATH) .


# ==================================================================================== #
# LOCAL TOOLS
# ==================================================================================== #

TOOLS_DIR ?= $(shell pwd)/.cache/tools

GOLANGCI_LINT_VERSION ?= v1.59.1
GOLANGCI_LINT = $(TOOLS_DIR)/golangci-lint-$(GOLANGCI_LINT_VERSION)

## golangci-lint: download golangci-lint locally if necessary.
golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT): $(TOOLS_DIR)
	$(call go-install-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint,${GOLANGCI_LINT_VERSION})

$(TOOLS_DIR):
	@mkdir -p $(TOOLS_DIR)

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(TOOLS_DIR) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef
