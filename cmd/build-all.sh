#!/bin/sh
cd ..
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./target/pixiv-proxy-go-win-amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./target/pixiv-proxy-go-linux-amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./target/pixiv-proxy-go-linux-arm64 main.go
