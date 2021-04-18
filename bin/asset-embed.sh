#!/bin/bash

## Embeds assets for building into the project

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "Embedding assets..."
go-embed -input assets -output app/assets/assets.go
