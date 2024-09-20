#!/bin/sh

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

export PATH="$PATH:/usr/local/go/bin"

go version

## Build step
# go install github.com/air-verse/air@latest
# export PATH=$PATH:$(go env GOPATH)/bin
# ls $(go env GOPATH)/bin/air
# cp /root/go/bin/air /app

make build
make run

./air -c air.toml