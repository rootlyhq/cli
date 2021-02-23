<div align="center">
  <img alt="logo" src="./docs/images/logo.png" height="250px">

  <h1>rootly cli</h1>

  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/rootly-io/cli">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/rootly-io/cli">
  <img alt="Golang report card" src ="https://goreportcard.com/badge/github.com/rootly-io/cli">
  <br>
  <img alt="build" src="https://github.com/rootly-io/cli/workflows/build/badge.svg" />
  <img alt="lint" src="https://github.com/rootly-io/cli/workflows/lint/badge.svg" />
  <img alt="release" src="https://github.com/rootly-io/cli/workflows/release/badge.svg" />
  <img alt="test" src="https://github.com/rootly-io/cli/workflows/test/badge.svg" />
  <br />
  <br />
  <i>A command-line interface for rootly</i>
</div>

<hr />

## 📜 Table of Contents

- [📜 Table of Contents](#-table-of-contents)
- [👋 Getting Started](#-getting-started)
  - [🚀 Install](#-install)
    - [🍎 macOS](#-macos)
    - [🐧 Linux and 🖥️ Windows](#-linux-and-️-windows)
- [📟 Commands](#-commands)
  - [ℹ️ `rootly pulse`](#ℹ️-rootly-pulse)
  - [🏃 `rootly pulse-run`](#-rootly-pulse-run)
- [📦 Running in CI](#-running-in-ci)

## 👋 Getting Started

The rootly command-line tool allows you to interact with rootly in your terminal, CI environment, or anything that can run a simple program. At the moment the main purpose of the cli is to send [pulses](https://rootly.io/docs/pulses) right from the command line. This is great for sending a pulse at the end of a deploy script for example. You can also send pulses based off the exit status of a given command to run.

### 🚀 Install

#### 🍎 macOS

Simply run the command below:

```bash
brew install rootly-io/homebrew-tap/rootly
```

#### 🐧 Linux and 🖥️ Windows

You can grab the binary from the [latest release](https://github.com/rootly-io/cli/releases/latest)

## 📟 Commands

### ℹ️ `rootly pulse`

`rootly pulse` allows you to send a [pulse](https://rootly.io/docs/pulses) right from the command-line. The summary for the pulse, which is required, goes at the end of the command as a normal argument.

| **Flag Name** | **Description**                                                         | **Examples**                           | **Required** |
| ------------- | ----------------------------------------------------------------------- | -------------------------------------- | ------------ |
| api-key       | A rootly api key                                                        | `--api-key "ABC123"`                   | Yes          |
| labels        | Key value pair labels (separated with commas with no spaces around `=`) | `--labels "Version=2, Attempt=1"`      | No           |
| services      | Services associated with the pulse (separated with commas)              | `--services "elasticsearch-prod"`      | No           |
| environments  | Environments associated with the pulse (separated with commas)          | `--environments "staging, production"` | No           |

Here are some examples:

- `rootly pulse-run --api-key "ABC123" Hello there`
  - Summary: Hello there
  - Labels: None
  - Services: None
  - Environments: None
- `rootly pulse --api-key "ABC123" --environments "staging, production" --services "elasticsearch-staging, elasticsearch-prod" Hello World!`
  - Summary: Hello World
  - Labels: None
  - Services: elasticsearch-staging and elasticsearch-prod
  - Environments: staging and production
- `rootly pulse --api-key ABC123 --environments "production" --labels "Version=2, Attempt=1" Is this thing on?`
  - Summary: Is this thing on?
  - Labels: Version: 2 and Attempt: 1
  - Services: None
  - Environments: production

### 🏃 `rootly pulse-run`

`rootly pulse-run` allows you to wrap a terminal command and send a [pulse](https://rootly.io/docs/pulses) with a label of the exit code. The summary for the pulse is a flag and if no value is provided it will use the command. The command goes at the end of the command as a normal argument.

| **Flag Name** | **Description**                                                         | **Examples**                           | **Required** |
| ------------- | ----------------------------------------------------------------------- | -------------------------------------- | ------------ |
| api-key       | A rootly api key                                                        | `--api-key "ABC123"`                   | Yes          |
| summary       | Summary for the pulse. Default is just the command                      | `--summary "Deployed Website"`         | No           |
| labels        | Key value pair labels (separated with commas with no spaces around `=`) | `--labels "Version=2, Attempt=1"`      | No           |
| services      | Services associated with the pulse (separated with commas)              | `--services "elasticsearch-prod"`      | No           |
| environments  | Environments associated with the pulse (separated with commas)          | `--environments "staging, production"` | No           |

Here are some examples:

- `rootly pulse-run --api-key "ABC123" make publish`
  - Summary: make publish
  - Labels: Exit Code: 0
  - Services: None
  - Environments: None
  - Command: `make publish`
- `rootly pulse-run --api-key "ABC123" --environments "staging, production" --services "elasticsearch-staging, elasticsearch-prod" echo Hello World`
  - Summary: echo Hello World
  - Labels: Exit Code: 0
  - Services: elasticsearch-staging and elasticsearch-prod
  - Environments: staging and production
  - Command: `echo Hello World`
- `rootly pulse --api-key ABC123 --environments "production" --labels "Version=2, Attempt=1" --summary "Deploy Website" sh deploy.sh`
  - Summary: sh deploy.sh
  - Labels: Version: 2, Attempt: 1, and Exit Code: 1
  - Services: None
  - Environments: production
  - Command: `sh deploy.sh`

## 📦 Running in CI

When using the rootly CLI in a CI environment there are some useful features to make the process easier. Every single flag can use a environment variable instead. The `api-key` flag for example could use the environment variable `ROOTLY_CLI_API_KEY` instead. To get the environment variable for a certain flag just replace all hyphens (`-`) with underscores (`_`), make all letters uppercase, and add `ROOTLY_CLI_` to the front.

There is also a GitHub action for `rootly pulse` that makes it really easy to use in a GitHub actions environment. See the [rootly-io/pulse-action](https://github.com/rootly-io/pulse-action) repository for more information.
