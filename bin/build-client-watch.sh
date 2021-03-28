#!/bin/bash

## Builds the TypeScript resources using `build-client`, then watches for changes
## Requires tsc available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../client

echo "Watching TypeScript compilation for [client/src]..."
esbuild src/client.ts --bundle --sourcemap --target=chrome58,firefox57,safari11,edge16 --outfile=../web/assets/client.js --watch
