#!/bin/bash
echo $CAPTUROO_VERSION
GOOS=linux   GOARCH=amd64 go build -ldflags "-X main.version=$CAPTUROO_VERSION" -o build/capturoo-$CAPTUROO_VERSION
GOOS=darwin  GOARCH=amd64 go build -ldflags "-X main.version=$CAPTUROO_VERSION" -o build/capturoo-$CAPTUROO_VERSION-darwin
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$CAPTUROO_VERSION" -o build/capturoo-$CAPTUROO_VERSION-windows
