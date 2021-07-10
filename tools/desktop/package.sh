#!/bin/bash

## Uses `tools/desktop` to build a desktop application

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

go-embed -input assets -output app/assets/assets.go

cd tools/desktop

go mod download

echo "starting macOS desktop build..."
GOOS=darwin GOARCH=amd64 CC=o64-clang CXX=o64-clang++ go build -o ../../dist/darwin_amd64/admini

echo "starting Linux desktop build..."
GOOS=linux GOARCH=amd64 go build -o ../../dist/linux_amd64/admini

echo "starting Windows desktop build..."
GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -o ../../dist/windows_amd64/admini
