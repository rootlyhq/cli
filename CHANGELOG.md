# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Add CLAUDE.md documentation for Claude Code assistance
- Add version management helpers and Makefile improvements
  - Add `scripts/bump-version.sh` for automated semantic versioning
  - Add new Makefile targets: `version-show`, `version-patch/minor/major`, `release-patch/minor/major`
  - Add help target with self-documenting commands

### Changed
- Update Go to 1.25 and upgrade all dependencies
  - github.com/spf13/cobra: v1.9.1 → v1.10.1
  - github.com/stretchr/testify: v1.10.0 → v1.11.1
  - github.com/oapi-codegen/oapi-codegen/v2: v2.4.1 → v2.5.1
  - github.com/getkin/kin-openapi: v0.132.0 → v0.133.0
  - golang.org/x/sys: v0.33.0 → v0.37.0
  - golang.org/x/text: v0.25.0 → v0.30.0

## [v1.2.10] - 2025-06-04

### Fixed
- Fix panic in key-value parsing for refs and labels
- Fix pulse command usage in help text

## [v1.2.9] - 2025-06-04

### Changed
- Update the version number
- Merge pull request #11 from er0k/fix-go-install-instructions

## [v1.2.8] - 2025-06-04

### Fixed
- Fix GoReleaser configuration for static binary builds

## [v1.2.7] - 2025-06-04

### Changed
- Update GitHub workflows to Go 1.24 and fix deprecated lint format
- Clean up dependencies with go mod tidy
- Fix linting issues
- Upgrade to Go 1.24 and fix static linking for binary compatibility
- Add .claude to gitignore
- Fix typo in README: correct ROOTLY_REFS environment variable

## [v1.2.6] - 2025-02-18

### Changed
- Upgrade to Go 1.23

## [v1.2.5] - 2025-02-18

### Changed
- Upgrade to Go 1.23

## [v1.2.4] - 2024-07-22

### Changed
- Upgrade dependencies

## [v1.2.3] - 2024-06-04

### Changed
- Upgrade to Go 1.22

## [v1.2.2] - 2024-06-04

### Changed
- Upgrade dependencies

## [v1.2.1] - 2024-02-09

### Changed
- Get values from ENV first (prioritize environment variables)

## [v1.2.0] - 2023-11-03

### Fixed
- Statically compile binary for better compatibility

## [v1.1.9] - 2023-11-03

### Fixed
- Use golang alpine image for Docker builds

## [v1.1.8] - 2023-11-02

### Changed
- Upgrade rootly-go and all other dependencies

## [v1.1.7] - 2022-05-17

### Fixed
- Run go mod tidy to clean up dependencies

## [v1.1.6] - 2021-10-17

### Fixed
- Run go mod tidy to clean up dependencies

## [v1.1.5] - 2021-09-18

### Added
- Add go install method of installation to documentation

## [v1.1.4] - 2021-08-07

### Changed
- Change organization from rootly-io to rootlyhq

## [v1.1.3] - 2021-07-29

### Added
- Add release target to scripts

### Changed
- Upgrade all dependencies

## [v1.1.1] - 2021-04-14

### Changed
- Upgrade all dependencies

## [v1.1.0] - 2021-04-10

### Added
- Add source and refs inputs for pulses
- Default source is now "cli"

### Changed
- Update GitHub action example documentation

## [v1.0.1] - 2021-02-24

### Added
- Add install script for easy installation

### Fixed
- Fix go report card badge
- Fix pulse run example in documentation
- Remove pulse-run from pulse example

## [v1.0.0] - 2021-02-24

### Added
- Add LICENSE file
- Add CONTRIBUTING documentation
- Add `--version` flag to check for updates
- Add debug and quiet/silent log modes
- Add complete pulse and pulse-run commands
- Add label support with key-value pairs
- Add environment and service flags
- Add API host configuration

### Changed
- Refactor entire config system
- Force key format in labels (lowercase with underscores)
- Upgrade to go1.16
- Switch to golangci-lint v2.5.0

### Fixed
- Fix multiple bug fixes and improvements
- Improve error handling for CreatePulse
- Handle failed environment IDs and service IDs properly

## [v0.0.4] - 2021-02-23

### Changed
- Run go mod tidy

### Fixed
- Small bug fixes and rootly-go upgrade
- Fix environment variable tests
- Add api-host flag
- Change environment variable prefix to ROOTLY_

## [v0.0.3] - 2021-02-22

### Added
- Add comprehensive documentation

### Fixed
- Fix goreleaser configuration
- Change binary goreleaser build name to rootly

## [v0.0.2] - 2021-02-22

### Changed
- Group array flags under one flag

### Fixed
- Properly handle failed environment IDs and service IDs
- Check for correct status code in requests

### Added
- Get summary if running in GitHub Action
- Improve flag descriptions
- Prep for GitHub Action integration

## [v0.0.1] - 2021-02-21

### Added
- Initial release
- Basic pulse command functionality
- Docker support
- CI/CD pipeline
- Unit tests
- Linting with golangci-lint and hadolint

[Unreleased]: https://github.com/rootlyhq/cli/compare/v1.2.10...HEAD
[v1.2.10]: https://github.com/rootlyhq/cli/compare/v1.2.9...v1.2.10
[v1.2.9]: https://github.com/rootlyhq/cli/compare/v1.2.8...v1.2.9
[v1.2.8]: https://github.com/rootlyhq/cli/compare/v1.2.7...v1.2.8
[v1.2.7]: https://github.com/rootlyhq/cli/compare/v1.2.6...v1.2.7
[v1.2.6]: https://github.com/rootlyhq/cli/compare/v1.2.5...v1.2.6
[v1.2.5]: https://github.com/rootlyhq/cli/compare/v1.2.4...v1.2.5
[v1.2.4]: https://github.com/rootlyhq/cli/compare/v1.2.3...v1.2.4
[v1.2.3]: https://github.com/rootlyhq/cli/compare/v1.2.2...v1.2.3
[v1.2.2]: https://github.com/rootlyhq/cli/compare/v1.2.1...v1.2.2
[v1.2.1]: https://github.com/rootlyhq/cli/compare/v1.2.0...v1.2.1
[v1.2.0]: https://github.com/rootlyhq/cli/compare/v1.1.9...v1.2.0
[v1.1.9]: https://github.com/rootlyhq/cli/compare/v1.1.8...v1.1.9
[v1.1.8]: https://github.com/rootlyhq/cli/compare/v1.1.7...v1.1.8
[v1.1.7]: https://github.com/rootlyhq/cli/compare/v1.1.6...v1.1.7
[v1.1.6]: https://github.com/rootlyhq/cli/compare/v1.1.5...v1.1.6
[v1.1.5]: https://github.com/rootlyhq/cli/compare/v1.1.4...v1.1.5
[v1.1.4]: https://github.com/rootlyhq/cli/compare/v1.1.3...v1.1.4
[v1.1.3]: https://github.com/rootlyhq/cli/compare/v1.1.1...v1.1.3
[v1.1.1]: https://github.com/rootlyhq/cli/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/rootlyhq/cli/compare/v1.0.1...v1.1.0
[v1.0.1]: https://github.com/rootlyhq/cli/compare/v1.0.0...v1.0.1
[v1.0.0]: https://github.com/rootlyhq/cli/compare/v0.0.4...v1.0.0
[v0.0.4]: https://github.com/rootlyhq/cli/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/rootlyhq/cli/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/rootlyhq/cli/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/rootlyhq/cli/releases/tag/v0.0.1
