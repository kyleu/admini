#!/bin/bash

## Runs goreleaser

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

goreleaser release --rm-dist
