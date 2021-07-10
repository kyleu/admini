#!/bin/bash

## Uses `tools/svgbuild` to generate Go code for the svgs in `client`

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

cd tools/svgbuilder

make build

cd $dir/../..
tmp/svgbuilder client/src/svg app/util/svg.go
