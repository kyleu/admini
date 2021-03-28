#!/bin/bash

## Uses `esbuild` to compile the scripts in `client`
## Requires esbuild and watchexec available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../client

esbuild src/client.ts --bundle --sourcemap --target=chrome58,firefox57,safari11,edge16 --outfile=../web/assets/client.js
