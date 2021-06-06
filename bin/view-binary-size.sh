#!/bin/bash

## Visualizes space usage for the release binary

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

go tool nm -size build/release/admini | c++filt > ./tmp/nm.txt

