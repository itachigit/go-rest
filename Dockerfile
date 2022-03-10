# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN apk add git
RUN apk add postgresql-client
RUN go mod download

COPY . .

RUN go build ./cmd/go-cowin/main.go

EXPOSE 8080