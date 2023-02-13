# Written using: https://docs.docker.com/language/golang/build-images/

FROM golang:1.19.5-alpine3.17
LABEL maintainer="Rohit Awate"

WORKDIR /app

# Copy all source files
COPY . /app

# Download necessary libraries
RUN go mod download

EXPOSE 8080

# Build and run the app
RUN go build -o cardhero
CMD ["./cardhero"]