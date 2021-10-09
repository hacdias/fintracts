# Fintracts English Parser

Fintracts parser for the english grammar specification. Translates an [english contract](SPECIFICATION.md) to the common [JSON format](../SPECIFICATION.md).

- [Run Directly](#run-directly)
- [Run With Docker](#run-with-docker)
- [License](#license)

## Run Directly

First, install the necessary dependencies:

```bash
go get ./...
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
