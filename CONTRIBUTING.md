# Contributing

Thank you for having an interest in contributing to the rootly CLI!

- [Contributing](#contributing)
  - [Setting up your environment](#setting-up-your-environment)
  - [Checking your code](#checking-your-code)
    - [Unit-tests](#unit-tests)
    - [Linting](#linting)
  - [Building](#building)
  - [Release](#release)

## Setting up your environment

1. Clone this repository
2. Run `go get all` to get all the dependencies needed.

## Checking your code

### Unit-tests

Just run the following command:

```bash
make test
```

### Linting

Make sure you have the following programs installed:

- [golangci-lint](https://github.com/golangci/golangci-lint)
- [hadolint](https://github.com/hadolint/hadolint)
- [goreleaser](https://github.com/goreleaser/goreleaser)

Once you have all the programs listed above run the following command:

```bash
make lint
```

## Building

Build a binary that will go in a folder called bin with the following command:

```bash
make build
```

You can also build the docker image with the following command:

```bash
make docker-build
```

## Release

You can release a new version of the rootly cli (if you have contributor access) by running the following commands:

```bash
make release VERSION="v1.0.0"
```

_Replace v1.0.0 with the version you want to release_
