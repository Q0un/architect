# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o users cmd/users/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /build/users ./
COPY ./configs/config.yaml ./
COPY ./tokens ./

CMD ./users 

EXPOSE 5000
