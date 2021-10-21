FROM golang:1.17-alpine AS build

WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download
COPY ../ .
RUN go build -o /validator validator/cmd/*.go

FROM alpine
WORKDIR /
COPY --from=build /validator /validator
ENTRYPOINT ["/validator"]
