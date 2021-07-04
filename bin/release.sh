#!/bin/bash

## Runs goreleaser

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

./bin/asset-embed.sh
git update-index --assume-unchanged app/assets/assets.go
goreleaser release --rm-dist
./bin/asset-reset.sh
git update-index --assume-unchanged app/assets/assets.go
