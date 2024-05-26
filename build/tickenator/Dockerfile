# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tickenator cmd/tickenator/main.go

FROM alpine:3.16

WORKDIR /app

RUN apk update && \
    apk add --no-cache \
    openssh-keygen \
    openssl

COPY --from=builder /build/tickenator ./
COPY ./configs/tickenator.yaml ./config.yaml

CMD ./tickenator

EXPOSE 5001
