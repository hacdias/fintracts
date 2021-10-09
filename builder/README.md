# Fintracts Builder

- [Run Directly](#run-directly)
- [Run With Docker](#run-with-docker)
- [License](#license)


## Run Directly

First, install the necessary dependencies:

```bash
npm ci
```

To start a development server at [localhost:3000](http://localhost:3000):

```bash
npm start
```

To build the app:

```bash
npm run build
```

The app will be built in a `dist/` directory. It contains plain HTML, CSS and JS files and it can be served by any web server.

## Run With Docker

The Docker image runs the development server and expects the source code volume to be mounted at `/app`. Dependencies should be automatically installed. If there's any issues, please remove `node_modules` before running the image.

```bash
# Build image
docker build . -t fintracts/builder

# Run image
docker run -it -p 3000:3000 -v $(pwd):/app fintracts/builder
```

You can also start a shell and run any other commands inside the container:

```bash
docker run -it -p 3000:3000 -v $(pwd):/app --entrypoint /bin/bash fintracts/builder
```

## License

[MIT © Henrique Dias](../LICENSE)