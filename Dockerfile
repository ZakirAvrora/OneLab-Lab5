FROM golang:1.17-alpine3.16 AS builder
LABEL stage=builder
WORKDIR /app

COPY . .

RUN go build -o main ./cmd/.
