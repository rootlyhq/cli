# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

The Rootly CLI is a command-line interface for interacting with rootly.com, primarily focused on sending "pulses" (deployment notifications and tracking events) from the command line, CI environments, or shell scripts. It's written in Go and uses the Cobra framework for command structure.

## Common Commands

### Development
```bash
# Install dependencies
go get all

# Build the binary (outputs to ./bin/rootly)
make build

# Run tests
make test

# Run linters (requires golangci-lint, hadolint, goreleaser)
make lint

# Clean build artifacts
make clean
```

### Testing Individual Components
```bash
# Run tests for a specific package
go test -count=1 -v ./pkg/api/...
go test -count=1 -v ./pkg/commands/...
go test -count=1 -v ./pkg/inputs/...
```

### Docker
```bash
# Build Docker image
make docker-build

# Push Docker image
make docker-push
```

### Release & Version Management

The project uses semantic versioning with automated version bumping:

```bash
# Show current and next versions
make version-show

# Bump versions (updates root.go, commits, creates tag, and pushes)
make release-patch  # 1.2.10 → 1.2.11 (for bug fixes)
make release-minor  # 1.2.10 → 1.3.0 (for new features)
make release-major  # 1.2.10 → 2.0.0 (for breaking changes)

# Or use individual steps
make version-patch  # Just bump and tag patch version
make version-minor  # Just bump and tag minor version
make version-major  # Just bump and tag major version

# Preview next version without making changes
make version-next

# Legacy manual release (requires VERSION variable)
make release VERSION="v1.0.0"
```

**Note**: Version bumping requires a clean git working directory. The script will:
1. Update version in `pkg/commands/root.go`
2. Commit the change
3. Create and push a git tag
4. Trigger GitHub Actions to build and publish the release

## Architecture

### Package Structure

- **cmd/rootly/** - Main entry point; initializes logging and executes root command
- **pkg/commands/** - Cobra command definitions (`pulse`, `pulse-run`, root command)
- **pkg/api/** - HTTP client logic for Rootly API, request building, and response parsing
- **pkg/inputs/** - Input handling across three layers:
  - `flags/` - Cobra flag definitions
  - `env/` - Environment variable parsing
  - `parse/` - String parsing utilities (arrays, key-value pairs)
  - Main package unifies flag + env variable retrieval
- **pkg/models/** - Data structures for pulses and API responses
- **pkg/log/** - Custom logging with context-aware error handling

### Input Processing Flow

The CLI uses a unified input system that prioritizes flags over environment variables:

1. **Flag Layer** (`pkg/inputs/flags/`): Defines all Cobra flags
2. **Environment Layer** (`pkg/inputs/env/`): Provides fallback via `ROOTLY_*` env vars
3. **Input Getter** (`pkg/inputs/get.go`): Unified retrieval that checks flags first, then env vars
4. **Parsing Layer** (`pkg/inputs/parse/`): Converts strings to structured data (arrays, maps)

Key pattern: `inputs.GetString(names.ApiKeyName, cmd, required)` checks flag `--api-key`, then falls back to `ROOTLY_API_KEY`.

### Commands

**`rootly pulse`** (`pkg/commands/pulse.go`)
- Sends a pulse with provided metadata
- Summary is required (taken from positional arguments or `ROOTLY_SUMMARY`)
- Collects metadata: labels (key=value pairs), services, environments, source, refs
- Timestamps: `StartedAt` captured at command start, `EndedAt` when pulse is sent

**`rootly pulse-run`** (`pkg/commands/pulse_run.go`)
- Wraps a shell command execution
- Automatically adds `exit_status` label with the command's exit code
- If no summary provided, uses the command itself as summary
- Flow: Parse inputs → Run command → Add exit code to labels → Send pulse

### API Integration

- Uses generated client from `rootlyhq/rootly-go` package
- Manual request construction in `pkg/api/request.go` using `oapi-codegen`
- Authentication via Bearer token (API key)
- Validates service/environment IDs in response and reports missing ones

### Key-Value Parsing

Labels and refs use format `key=value, key2=value2` (comma-separated, no spaces around `=`).

Parser (`pkg/inputs/convert.go`):
- Splits on commas, then on `=`
- Converts keys to lowercase and replaces spaces with underscores
- Validates format (exactly one `=` per item, non-empty values)

Critical fix context: Recent fix addressed panic in key-value parsing for refs and labels when format was invalid (see commit 2c38ee2).

## Version Management

Current version is maintained in `pkg/commands/root.go:34` and managed through git tags.

### Version Bumping Process

The `scripts/bump-version.sh` script automates version management:
1. Reads current version from latest git tag
2. Calculates next version based on semver (patch/minor/major)
3. Updates `pkg/commands/root.go` with new version
4. Commits the change with message "Update version to vX.Y.Z"
5. Creates annotated git tag (vX.Y.Z)
6. Pushes commit and tag to origin

### Usage

Use Makefile targets for version management:
- `make version-show` - Display current and next versions
- `make release-patch` - Bump patch (bug fixes)
- `make release-minor` - Bump minor (new features)
- `make release-major` - Bump major (breaking changes)

The CLI checks for updates on `--version` flag using the `gleich/release` package.

## CI/CD

GitHub Actions workflows (`.github/workflows/`):
- `build.yml` - Build verification
- `lint.yml` - Linting checks
- `test.yml` - Test execution
- `release.yml` - Automated releases via GoReleaser

All flags support environment variables: replace hyphens with underscores, uppercase, and prefix with `ROOTLY_`.

## Testing Notes

- Tests use `testify` assertions
- Example test locations:
  - API conversion tests: `pkg/api/convert_test.go`
  - Request tests: `pkg/api/request_test.go`
  - Parsing tests: `pkg/inputs/parse/array_test.go`
  - Logging tests: `pkg/log/format_test.go`
