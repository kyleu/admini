#!/bin/bash

## Makes a release build, builds a docker image, then exports and zips the output
## Requires docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

docker build -m 4g -t kyleu/admini .

mkdir -p build/docker
docker save -o build/docker/admini.docker.tar kyleu/admini
cd build/docker/
rm -f admini.docker.tar.gz
gzip admini.docker.tar
