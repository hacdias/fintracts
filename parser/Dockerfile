FROM golang:1.17-alpine AS build

WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download
COPY ../ .
RUN go build -o /parser parser/cmd/*.go

FROM alpine
WORKDIR /
COPY --from=build /parser /parser
ENTRYPOINT ["/parser"]
