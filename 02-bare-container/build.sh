#!/bin/bash

# uncomment to cross build linux binary from macOS/Windows
# export GOOS=linux
# export GOARCH=amd64

go build -o server server.go

docker build -t bare-go .
