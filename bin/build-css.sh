#!/bin/bash

## Uses `scss` to compile the stylesheets in `web/stylesheets`
## Requires SCSS available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

sass --no-source-map web/stylesheets/client.scss web/assets/client.css
sass --style=compressed --no-source-map web/stylesheets/client.scss web/assets/client.min.css
