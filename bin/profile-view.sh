#!/bin/bash

## Runs code statistics, checks for outdated dependencies, then runs linters

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

echo "=== launching profiler for cpu.pprof ==="
go tool pprof -http=":8000" ./build/debug/admini ./cpu.pprof

