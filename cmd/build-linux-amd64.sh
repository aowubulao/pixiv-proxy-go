#!/bin/sh
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./target/pixiv-proxy-go-linux-amd64 main.go
