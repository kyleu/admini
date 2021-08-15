#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

## Uses `esbuild` to compile the scripts in `client`
## Requires esbuild and watchexec available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../../client

node build.js