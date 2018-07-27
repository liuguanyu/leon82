#!/bin/bash

# GOFLAGS='-ldflags="-s -w"'
version=`git describe --tags`
arch=$(go env GOARCH)

for os in linux darwin windows; do
    echo "... building $version for $os/$arch"
    TARGET="leon82-$os-$arch"
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build
    if [ "$os" == "windows" ]; then
        tar czvf $TARGET.tar.gz views leon82.exe
    else
        tar czvf $TARGET.tar.gz views leon82
    fi
done