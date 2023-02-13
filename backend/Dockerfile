# syntax=docker/dockerfile:1
# Written using: https://docs.docker.com/language/golang/build-images/

## Build
FROM golang:1.19.5-alpine3.17 as build
LABEL maintainer="Rohit Awate"

WORKDIR /app
# Copy all source files
COPY . /app

# Download necessary libraries
RUN go mod download

# Build and run the app
RUN go build -o cardhero

## Deploy
FROM gcr.io/distroless/base-debian10:latest as deploy

WORKDIR /app

COPY --from=build /app/cardhero /cardhero

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/cardhero"]