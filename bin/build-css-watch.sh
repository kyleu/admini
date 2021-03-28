#!/bin/bash

## Builds the css resources using `build-css`, then watches for changes in `stylesheets`
## Requires SCSS available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

bin/build-css.sh
echo "Watching sass compilation for [web/stylesheets/client.scss]..."
sass --watch --no-source-map web/stylesheets/client.scss web/assets/client.css
