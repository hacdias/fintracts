FROM ocaml/opam:alpine-3.14 AS base

# Update and upgrade default packages
RUN sudo apk update && \
  sudo apk upgrade && \
  sudo apk add gmp-dev musl-dev linux-headers

WORKDIR /app

# Install dependencies
COPY *.opam .
COPY Makefile .
RUN make install-deps

# Build application
COPY . .
RUN eval $(opam env) && \
  make build && \
  sudo cp ./fintracts.exe /usr/bin/fintracts.exe

FROM alpine:3.14

COPY --from=base /usr/bin/fintracts.exe /home/app/fintracts.exe

RUN adduser -D app && \
  chown app:app /home/app/fintracts.exe

WORKDIR /home/app
USER app
ENTRYPOINT ["/home/app/fintracts.exe"]
