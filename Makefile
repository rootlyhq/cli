.PHONY: build docker-build docker-push clean test lint release help
.PHONY: version-show version-patch version-minor version-major version-next version-help
.PHONY: release-patch release-minor release-major

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

build: ## Build the binary
	CGO_ENABLED=0 go build -tags netgo -a -v -ldflags '-extldflags "-static"' -o ./bin/rootly ./cmd/rootly/

docker-build: ## Build Docker image
	docker build -t rootlyhub/cli .

docker-push: ## Push Docker image
	docker push rootlyhub/cli

clean: ## Remove build artifacts
	rm -r ./bin

test: ## Run tests
	go test -count=1 -v ./...

lint: ## Run linters
	golangci-lint run
	hadolint Dockerfile
	goreleaser check

release: ## Create a release with VERSION variable (legacy)
	git tag -a $(VERSION) -m $(VERSION)
	git push origin $(VERSION)

# Version management targets
# These targets manage semantic versioning using git tags

version-show: ## Show current and next versions
	@echo "Current version: $$(git describe --tags --abbrev=0 2>/dev/null || echo 'No tags found')"
	@echo "Next patch: $$(scripts/bump-version.sh show patch)"
	@echo "Next minor: $$(scripts/bump-version.sh show minor)"
	@echo "Next major: $$(scripts/bump-version.sh show major)"

version-patch: ## Bump patch version (1.2.3 → 1.2.4)
	@scripts/bump-version.sh patch

version-minor: ## Bump minor version (1.2.3 → 1.3.0)
	@scripts/bump-version.sh minor

version-major: ## Bump major version (1.2.3 → 2.0.0)
	@scripts/bump-version.sh major

version-next: ## Show next patch version
	@scripts/bump-version.sh show patch

version-help: ## Show detailed version help
	@scripts/bump-version.sh help

# Release targets - these create git tags which trigger CI releases

release-patch: version-patch ## Bump patch version and push tag (triggers CI release)
	@echo "✅ Patch version bumped and tagged"
	@echo "🚀 GitHub Actions will automatically build and publish the release"

release-minor: version-minor ## Bump minor version and push tag (triggers CI release)
	@echo "✅ Minor version bumped and tagged"
	@echo "🚀 GitHub Actions will automatically build and publish the release"

release-major: version-major ## Bump major version and push tag (triggers CI release)
	@echo "✅ Major version bumped and tagged"
	@echo "🚀 GitHub Actions will automatically build and publish the release"
