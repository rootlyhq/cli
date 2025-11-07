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

### Release
```bash
# Create and push a new version tag
make release VERSION="v1.0.0"
```

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

Current version is hardcoded in `pkg/commands/root.go:34`. Update this file when releasing new versions.

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
