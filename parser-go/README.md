# Fintracts Parser

- [Run Directly](#run-directly)
- [Known Limitations](#known-limitations)
- [License](#license)


## Run Directly

<!-- First, install the necessary dependencies:

```bash
make install-deps
```

Then, build the executable:

```bash
make build
```

An executable can be found on `./fintracts.exe`, which can be executed:

```
./fintracts.exe < ./path/to/contract.txt
``` -->

## Run With Docker
<!-- 
To build:

```bash
docker build . -t fintracts/parser
```

You can now run the parser inside the docker image:

```bash
docker run -i fintracts/parser < ./path/to/contract.txt
``` -->

## Known Limitations

<!-- - The contract text is case-insensitive, i.e., "The parties" and "tHe PaRtIeS" are both valid.
- Option does not support special characters yet, i.e., "USD-LIBOR" is not valid yet. -->

## License

[MIT © Henrique Dias](../LICENSE)