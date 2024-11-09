#!/bin/sh

# export MYSQL_URL=$(busybox ping -c 1 mysql | busybox head -1 | busybox awk -F'[()]' '{print $2}')

rm -rf /usr/local/go && tar -C /usr/local -xzf src/go1.22.5.linux-amd64.tar.gz

export PATH="$PATH:/usr/local/go/bin"

# go get -u github.com/gin-gonic/gin
# go get -u github.com/go-gorp/gorp
# go get -u github.com/go-sql-driver/mysql

go version

cd src

# echo "go mod download ... start"

go mod download

# echo "go mod download ... done"

# go run main.go
go build -o hefa

echo "go build ... done"
ls

./hefa