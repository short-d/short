FROM golang:1.13.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git bash

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY . .

RUN go build -o build/app main.go

FROM alpine:3.10 as production

WORKDIR /app

COPY --from=builder /app/build/app ./build/app
COPY --from=builder /app/scripts ./scripts