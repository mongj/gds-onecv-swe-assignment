# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

WORKDIR /usr/src/server

ENV CGO_ENABLED 0
ENV GOPATH /go

COPY go.mod go.sum ./
RUN  go mod download

COPY . .
RUN go build -o bin/server main.go

RUN cp /usr/src/server/certificate.crt /etc/ssl/certificate.crt
RUN cp /usr/src/server/private.key /etc/ssl/private.key
RUN cp /usr/src/server/bin/server /usr/local/bin/server

EXPOSE 8000
ENTRYPOINT ["/usr/local/bin/server"]