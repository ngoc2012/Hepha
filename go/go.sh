#!/bin/sh

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

export PATH="$PATH:/usr/local/go/bin"

go version

go mod download

make build

make run

go install github.com/air-verse/air@latest

air -c air.toml