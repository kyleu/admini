#!/bin/bash

## Packages the build output for GitHub Releases

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..
pdir="$( pwd )"

# ./bin/build-all.sh

mkdir -p ./build/stage

rm -rf ./build/package
mkdir -p ./build/package

cd "$pdir/build/stage"

# cp -R "$pdir/data" ./data

pkg () {
  echo "$4 ($2)..."
  cp "$pdir/build/$1/$2/$3" "./$3"

  if [ $2 = "amd64" ]; then
    zip -r "$pdir/build/package/admini.$4.zip" *
  else
    zip -r "$pdir/build/package/admini.$4.$2.zip" *
  fi

  rm "./$3"
}

# macOS
pkg darwin amd64 admini macos
pkg darwin arm64 admini macos

# Linux
pkg linux amd64 admini linux
pkg linux 386 admini linux
pkg linux arm64 admini linux
pkg linux arm admini linux
# pkg linux mips admini linux
# pkg linux riscv64 admini linux

# FreeBSD
# pkg freebsd amd64 admini freebsd
# pkg freebsd 386 admini freebsd
# pkg freebsd arm64 admini freebsd
# pkg freebsd arm admini freebsd

# Windows
pkg windows amd64 admini.exe windows
pkg windows 386 admini.exe windows
# pkg windows arm admini.exe windows

# Docker
echo "docker..."
cp "$pdir/build/docker/admini.docker.tar.gz" "$pdir/build/package/admini.docker.tar.gz"

rm -rf "$pdir/build/stage"
