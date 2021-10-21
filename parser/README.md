# Fintracts English Parser

[![Build](https://img.shields.io/github/workflow/status/hacdias/fintracts/ci?style=flat-square)](https://github.com/hacdias/fintracts/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacdias/fintracts?style=flat-square)](https://goreportcard.com/report/github.com/hacdias/fintracts)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/hacdias/fintracts)

Fintracts parses contracts in the [English specification](./SPECIFICATION.md) to the common [JSON format](../SPECIFICATION.md).

- [Run Directly](#run-directly)
- [Usage](#usage)
- [Docker](#docker)
- [License](#license)


## Run Directly

First, install the necessary dependencies:

```bash
go mod download
```

Then, build the executable:

```bash
make build
```

An executable can be found on `./parser`.

## Usage

```bash
./parser < ./path/to/contract.txt
```

Run `./parser --help` for more information.

## Docker

To build the Docker image:

```bash
make docker
```

You can now run the validator inside the docker image:

```bash
docker run -i parser [commands and arguments]
```

Usage example:

```bash
docker run -i parser < ./path/to/contract.txt
```

## License

[MIT © Henrique Dias](../LICENSE)
