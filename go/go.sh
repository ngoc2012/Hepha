#!/bin/sh

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

export PATH="$PATH:/usr/local/go/bin"

go version

go mod download

ls
echo pwd

make build

make run

./run