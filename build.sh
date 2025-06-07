#!/bin/bash -ex

cd $(dirname $0)
rm -rf build
mkdir -p build

cd server
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o homedash_x86 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o homedash_arm .
mv homedash_* ../build
cp script/* ../build

cd ../ui
npm run build
mv dist ../build

cd ../build
tar czvf ~/homedash_pkg.tgz *
cd
sftp pi4 <<< 'put homedash_pkg.tgz'
ssh pi4 <<< 'pkill homedash_arm; rm -rf homedash_pkg && mkdir -p homedash_pkg && tar xzf homedash_pkg.tgz -C homedash_pkg'
