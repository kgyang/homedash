#!/bin/bash -ex

cd $(dirname $0)
mkdir -p build

cd server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o homedash .
#CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o homedash .
mv homedash ../build

cd ../ui
npm run build
mv dist ../build
