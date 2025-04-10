#!/usr/bin/env bash

set -ex

CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o build/windows/amd64/quick-echarts.exe
CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o build/linux/amd64/quick-echarts
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o build/macos/amd64/quick-echarts
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o build/macos/arm64/quick-echarts