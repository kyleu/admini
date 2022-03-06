# Content managed by Project Forge, see [projectforge.md] for details.
.PHONY: clean
clean: ## Removes builds and compiled templates
	@rm -rf tmp/*.hashcode
	@rm -rf out
	@find ./queries -type f -name '*.sql.go' -exec rm {} +
	@find ./views -type f -name '*.html.go' -exec rm {} +

.PHONY: dev
dev: ## Start the project, reloading on changes
	@bash bin/dev.sh

.PHONY: templates
templates:
	@bin/templates.sh

.PHONY: build
build: templates ## Build all binaries
	@go1.18rc1 build -gcflags "all=-N -l" -o build/debug/admini .

.PHONY: build-release
build-release: templates ## Build all binaries without debug information, clean up after
	@go1.18rc1 build -ldflags '-s -w' -trimpath -o build/release/admini .

.PHONY: lint
lint: ## Run linter
	@bin/check.sh

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
