# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o stats cmd/stats/main.go

FROM alpine:3.16

WORKDIR /app

RUN apk update && \
    apk add --no-cache \
    openssh-keygen \
    openssl

COPY --from=builder /build/stats ./
COPY ./configs/stats.yaml ./config.yaml

CMD ./stats

EXPOSE 5006
