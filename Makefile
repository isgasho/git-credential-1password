ifeq ($(OS),Windows_NT)
	SUFFIX := .exe
endif

ifndef VERBOSE
.SILENT: # no need for @
endif

# strip symbols
GO_BUILD_FLAGS := -s -w -extldflags "-static"
# go build command
GO_BUILD := CGO_ENABLED=0 go build
# go source files
GO_FILES = $(shell find . -name '*.go')

.DEFAULT_GOAL:=help

##@ Build

.PHONY: git-credential-1password

git-credential-1password: $(GO_FILES) ## Build git-credential-1password
	$(GO_BUILD) -ldflags '$(GO_BUILD_FLAGS)' -o bin/git-credential-1password$(SUFFIX) github.com/develerik/git-credential-1password

##@ Helpers

.PHONY: help

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <command> \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
