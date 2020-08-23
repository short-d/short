# FROM golang:1.13.1-alpine AS production

# RUN apk add --no-cache git bash

# COPY serve/go.mod serve/go.sum ./

# RUN go mod download

# RUN go mod verify

# COPY serve .

# COPY ./build ./public

# CMD 

# ENTRYPOINT ./scripts/start


FROM golang:1.13.1-alpine AS builder

WORKDIR /short

RUN apk add --no-cache git bash

# Install dependencies
COPY serve/go.mod serve/go.sum ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY serve .

RUN go build -o app main.go

FROM alpine:3.10 as production

WORKDIR /short

RUN apk add --no-cache bash

COPY ./build ./public

COPY --from=builder /short/app ./app

ENTRYPOINT ./app