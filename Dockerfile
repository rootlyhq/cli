FROM golang:1.25 AS builder

# Meta data
LABEL maintainer="support@rootly.com"
LABEL description="Command-line tool for rootly"

# Copying over all the files
COPY . /usr/src/app
WORKDIR /usr/src/app
# Installing dependencies and build the binary
RUN go mod download && \
    make build

# hadolint ignore=DL3006,DL3007
FROM alpine:latest
COPY --from=builder /usr/src/app/bin/rootly /usr/local/bin/rootly
