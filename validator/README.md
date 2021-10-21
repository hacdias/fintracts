# Fintracts Validator

[![Build](https://img.shields.io/github/workflow/status/hacdias/fintracts/ci?style=flat-square)](https://github.com/hacdias/fintracts/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacdias/fintracts?style=flat-square)](https://goreportcard.com/report/github.com/hacdias/fintracts)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/hacdias/fintracts)

Fintracts Validator semantically validates any JSON contract according to the [specification](../SPECIFICATION.md).

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

An executable can be found on `./validator`.

## Usage

```bash
./validator < ./path/to/contract.json
```

Run `./validator --help` for more information.

## Docker

Please note that to build this with Docker, we will need the entire repository. As such, you need to execute the following command to build:

```bash
docker build .. -t validator -f Dockerfile
```

You can now run the validator inside the docker image:

```bash
docker run -i validator [commands and arguments]
```

Usage example:

```bash
docker run -i validator < ./path/to/contract.json
```

## License

[MIT © Henrique Dias](../LICENSE)
