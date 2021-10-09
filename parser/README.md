# Fintracts English Parser

[![Build](https://img.shields.io/github/workflow/status/hacdias/fintracts/ci?style=flat-square)](https://github.com/hacdias/fintracts/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/hacdias/fintracts/parser?style=flat-square)](https://goreportcard.com/report/github.com/hacdias/fintracts/parser)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/hacdias/fintracts/parser)

Fintracts parser for the english grammar specification. Translates an [english contract](../ENGLISH-SPECIFICATION.md) to the common [JSON format](../JSON-SPECIFICATION.md).

- [Run Directly](#run-directly)
- [Run With Docker](#run-with-docker)
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

An executable can be found on `./parser`, which can be executed:

```
./parser < ./path/to/contract.txt
```

## Run With Docker

To build:

```bash
docker build . -t fintracts/parser
```

You can now run the parser inside the docker image:

```bash
docker run -i fintracts/parser < ./path/to/contract.txt
```

## License

[MIT © Henrique Dias](../LICENSE)
