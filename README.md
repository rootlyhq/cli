<div align="center">
  <img alt="logo" src="./docs/images/logo.png" height="250px">

  <h1>rootly cli</h1>

  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/rootlyhq/cli">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/rootlyhq/cli">
  <img alt="Golang report card" src="https://goreportcard.com/badge/github.com/rootlyhq/cli">
  <br>
  <img alt="build" src="https://github.com/rootlyhq/cli/workflows/build/badge.svg" />
  <img alt="lint" src="https://github.com/rootlyhq/cli/workflows/lint/badge.svg" />
  <img alt="release" src="https://github.com/rootlyhq/cli/workflows/release/badge.svg" />
  <img alt="test" src="https://github.com/rootlyhq/cli/workflows/test/badge.svg" />
  <br />
  <br />
  <i>A command-line interface for rootly</i>
</div>

<hr />

- [👋 Getting Started](#-getting-started)
  - [🚀 Install](#-install)
    - [🍎 macOS](#-macos)
    - [🐧 Linux and 🖥️ Windows](#-linux-and-️-windows)
- [📟 Commands](#-commands)
  - [ℹ️ `rootly pulse`](#ℹ️-rootly-pulse)
  - [🏃 `rootly pulse-run`](#-rootly-pulse-run)
- [📦 Running in CI](#-running-in-ci)
- [🤖 GitHub Action](#-github-action)

## 👋 Getting Started

The rootly command-line tool allows you to interact with rootly in your terminal, CI environment, or anything that can run a simple program. At the moment the main purpose of the cli is to send [pulses](https://rootly.io/docs/pulses) right from the command line. This is great for sending a pulse at the end of a deploy script for example. You can also send pulses based on the exit status of a given command to run.

### 🚀 Install

#### 🍎 macOS

Simply run the command below:

```bash
brew install rootlyhq/homebrew-tap/rootly
```

#### 🐧 Linux and 🖥️ Windows

You can grab the binary from the [latest release](https://github.com/rootlyhq/cli/releases/latest) or run the command below.

```bash
curl -fsSL https://raw.githubusercontent.com/rootlyhq/cli/main/install.sh | sh
```

If you have golang installed you can also just install the binary using the following command. Keep in mind that this will go in `~/go/bin/` by default which might not be in your path.

```bash
go install github.com/rootlyhq/cli@latest
```

## 📟 Commands

### ℹ️ `rootly pulse`

`rootly pulse` allows you to send a [pulse](https://rootly.io/docs/pulses) right from the command-line. The summary for the pulse, which is required, goes at the end of the command as a normal argument.

| **Flag Name** | **Description**                                                                                                                             | **Examples**                                            | **Required** | **Environment Variable** |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ------------ | ------------------------ |
| api-key       | A rootly api key                                                                                                                            | `--api-key "ABC123"`                                    | Yes          | ROOTLY_API_KEY           |
| api-host      | Host URL for the rootly api. Default is https://api.rootly.io                                                                               | `--api-host "https://rootly.hello.com"`                 | No           | ROOTLY_API_HOST          |
| labels        | Key value pair labels (separated with commas with no spaces around `=`)                                                                     | `--labels "version=2, attempt=1"`                       | No           | ROOTLY_LABELS            |
| services      | Services associated with the pulse (separated with commas)                                                                                  | `--services "elasticsearch-prod"`                       | No           | ROOTLY_SERVICES          |
| environments  | Environments associated with the pulse (separated with commas)                                                                              | `--environments "staging, production"`                  | No           | ROOTLY_ENVIRONMENTS      |
| sources       | Source of the pulse                                                                                                                         | `--source "k8s"`                                        | No           | ROOTLY_SOURCE            |
| refs          | Refs associated with the pulse. Separate each item with a comma. Key-value pair separated with an equal sign (= with no surrounding spaces) | `--refs "sha=cd62148cbc5eb42168fe99fdb50a364e12b206ac"` | No           | ROOTLY_ENVIRONMENTS      |
| debug         | If extra logs should be outputted for debugging                                                                                             | `--debug`                                               | No           | ROOTLY_DEBUG             |
| quiet         | If all logs should be suppressed                                                                                                            | `--quiet`                                               | No           | ROOTLY_QUIET             |

Here are some examples:

- `rootly pulse --api-key "ABC123" --quiet Hello there`
  - Summary: Hello there
  - Labels: None
  - Services: None
  - Environments: None
  - Source: cli
  - Refs None
  - Output: No logs
- `rootly pulse --api-key "ABC123" --environments "staging, production" --services "elasticsearch-staging, elasticsearch-prod" Hello World!`
  - Summary: Hello World
  - Labels: None
  - Services: elasticsearch-staging and elasticsearch-prod
  - Environments: staging and production
  - Source: cli
  - Refs None
  - Output: Silent
- `rootly pulse --api-key ABC123 --environments "production" --quiet --labels "Version=2, Attempt=1" --source "k8s" --refs "sha=cd62148cbc5eb42168fe99fdb50a364e12b206ac" Is this thing on?`
  - Summary: Is this thing on?
  - Labels: Version: 2 and Attempt: 1
  - Services: None
  - Environments: production
  - Source: k8s
  - Refs: sha: cd62148cbc5eb42168fe99fdb50a364e12b206ac
  - Output: No logs

### 🏃 `rootly pulse-run`

`rootly pulse-run` allows you to wrap a terminal command and send a [pulse](https://rootly.io/docs/pulses) with a label of the exit code. The summary for the pulse is a flag and if no value is provided it will use the command. The command goes at the end of the command as a normal argument.

| **Flag Name** | **Description**                                                                                                                             | **Examples**                                            | **Required** | **Environment Variable** |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ------------ | ------------------------ |
| api-key       | A rootly api key                                                                                                                            | `--api-key "ABC123"`                                    | Yes          | ROOTLY_API_KEY           |
| api-host      | Host URL for the rootly api. Default is https://api.rootly.io                                                                               | `--api-host "https://rootly.hello.com"`                 | No           | ROOTLY_API_HOST          |
| summary       | Summary for the pulse. Default is just the command                                                                                          | `--summary "Deployed Website"`                          | No           | ROOTLY_SUMMARY           |
| labels        | Key value pair labels (separated with commas with no spaces around `=`)                                                                     | `--labels "Version=2, Attempt=1"`                       | No           | ROOTLY_LABELS            |
| services      | Services associated with the pulse (separated with commas)                                                                                  | `--services "elasticsearch-prod"`                       | No           | ROOTLY_SERVICES          |
| environments  | Environments associated with the pulse (separated with commas)                                                                              | `--environments "staging, production"`                  | No           | ROOTLY_ENVIRONMENTS      |
| sources       | Source of the pulse                                                                                                                         | `--source "k8s"`                                        | No           | ROOTLY_SOURCE            |
| refs          | Refs associated with the pulse. Separate each item with a comma. Key-value pair separated with an equal sign (= with no surrounding spaces) | `--refs "sha=cd62148cbc5eb42168fe99fdb50a364e12b206ac"` | No           | ROOTLY_ENVIRONMENTS      |
| debug         | If extra logs should be outputted for debugging                                                                                             | `--debug`                                               | No           | ROOTLY_DEBUG             |
| quiet         | If all logs should be suppressed (command will still output)                                                                                | `--quiet`                                               | No           | ROOTLY_QUIET             |

Here are some examples:

- `rootly pulse-run --api-key "ABC123" make publish`
  - Summary: make publish
  - Labels: Exit Code: 0
  - Services: None
  - Environments: None
  - Source: cli
  - Refs: None
  - Command: `make publish`
  - Output: Regular logs
- `rootly pulse-run --quiet --api-key "ABC123" --environments "staging, production" --services "elasticsearch-staging, elasticsearch-prod" --source "k8s" echo Hello World`
  - Summary: echo Hello World
  - Labels: Exit Code: 0
  - Services: elasticsearch-staging and elasticsearch-prod
  - Environments: staging and production
  - Source: k8s
  - Refs: None
  - Command: `echo Hello World`
  - Output: No logs
- `rootly pulse-run --api-key "ABC123" --environments "production" --labels "version=2, attempt=1" --summary "Deploy Website" --source "k8s" --refs "sha=cd62148cbc5eb42168fe99fdb50a364e12b206ac" sh deploy.sh`
  - Summary: sh deploy.sh
  - Labels: Version: 2, Attempt: 1, and Exit Code: 1
  - Services: None
  - Environments: production
  - Source: k8s
  - Refs: sha: cd62148cbc5eb42168fe99fdb50a364e12b206ac
  - Command: `sh deploy.sh`
  - Output: Regular logs

## 📦 Running in CI

When using the rootly CLI in a CI environment there are some useful features to make the process easier. Every single flag can use an environment variable instead. The `api-key` flag for example could use the environment variable `ROOTLY_API_KEY` instead. To get the environment variable for a certain flag just replace all hyphens (`-`) with underscores (`_`), make all letters uppercase, and add `ROOTLY_` to the front.

## 🤖 GitHub Action

There is also a GitHub action for `rootly pulse` that makes it easy to use in a GitHub actions environment. See the [rootlyhq/pulse-action](https://github.com/rootlyhq/pulse-action) repository for more information. Here is a short little example:

```yaml
name: Deploy Website

on: push

jobs:
  pulse:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - run: make deploy
      - name: rootly-pulse
        uses: rootlyhq/pulse-action@main
        with:
          api_key: ${{ secrets.ROOTLY_API_KEY }}
          summary: Deploy Website
          environments: production # Not required
          services: elasticsearch-prod # Not required
          labels: platform=ubuntu,version=2 # Not required
          source: k8s # Not required
          refs: sha=cd62148cbc5eb42168fe99fdb50a364e12b206ac, image=registry.rootly.io/rootly/my-service:cd6214 # Not required
```
