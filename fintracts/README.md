# Fintracts CLI Tool

[![Build](https://img.shields.io/github/workflow/status/hacdias/fintracts/ci?style=flat-square)](https://github.com/hacdias/fintracts/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacdias/fintracts/fintracts?style=flat-square)](https://goreportcard.com/report/github.com/hacdias/fintracts/fintracts)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/hacdias/fintracts/fintracts)

Fintracts CLI tool parses English contracts and validates any JSON contract according to the specification.

- [Features](#features)
- [Install With Go](#install-with-go)
- [Run Directly](#run-directly)
- [Run With Docker](#run-with-docker)
- [Usage](#usage)
- [License](#license)

## Features

- Translates an [english contract](./english/SPECIFICATION.md) to the common [JSON format](../SPECIFICATION.md).
- Validates contracts in the common [JSON format](../SPECIFICATION.md).

## Install With Go

If you have Go installed, you can simply run:

```bash
go install github.com/hacdias/fintracts/fintracts/cmd/fintracts
```

## Run Directly

First, install the necessary dependencies:

```bash
go mod download
```

Then, build the executable:

```bash
make build
```

An executable can be found on `./fintracts`.

## Run With Docker

To build:

```bash
docker build . -t fintracts/parser
```

You can now run the parser inside the docker image:

```bash
docker run -i fintracts/parser [commands and arguments]
```

## Usage

Run `fintracts --help` for more information.

## License

[MIT © Henrique Dias](../LICENSE)
