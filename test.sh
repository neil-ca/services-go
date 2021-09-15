#! /bin/sh

starttest() {
	set -e
	GO111MODULE=on go test -race ./...
}

