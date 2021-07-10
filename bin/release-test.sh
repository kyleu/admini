#!/bin/bash

## Runs goreleaser in test mode

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

[[ -f "$HOME/bin/oauth" ]] && . $HOME/bin/oauth

./bin/asset-embed.sh
goreleaser --snapshot --skip-publish --rm-dist
./bin/asset-reset.sh
