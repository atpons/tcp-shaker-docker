FROM golang:latest as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/atpons/tcp-shaker-docker
COPY . .
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o tcp-shaker-cli main.go

FROM frolvlad/alpine-glibc

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/atpons/tcp-shaker-docker/tcp-shaker-cli /usr/local/bin/tcp-shaker-cli

ENTRYPOINT ["/usr/local/bin/tcp-shaker-cli"]