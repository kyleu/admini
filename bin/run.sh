#!/bin/bash

## Compiles and starts the server

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "running [admini]"
make build
./build/admini server admini
