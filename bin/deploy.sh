#!/bin/bash

## Deploys the app to my tiny server

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

GOOS=linux GOARCH=amd64 make build-release
mkdir -p ./build/linux/amd64
mv ./build/release/admini ./build/linux/amd64/admini
../kyleu.dev/deploy/admini.sh
