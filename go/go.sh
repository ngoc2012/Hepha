#!/bin/sh

rm -rf /usr/local/go && tar -C /usr/local -xzf src/go1.22.5.linux-amd64.tar.gz

export PATH="$PATH:/usr/local/go/bin"

go version

go mod download

make build

make run

./run