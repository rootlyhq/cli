FROM golang:1.15 AS builder

# Meta data
LABEL maintainer="support@rootly.io"
LABEL description="Command-line tool for rootly"

# Copying over all the files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies
RUN go get -v -t -d ./...

# Build the binary
RUN make build

# hadolint ignore=DL3006,DL3007
FROM alpine:latest
COPY --from=builder /usr/src/app/bin/rootly .
