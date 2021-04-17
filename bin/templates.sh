#!/bin/bash

## Builds all the templates using quicktemplate, skipping if unchanged

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

FORCE="${1:-}"

function tmpl {
  echo "updating [$2] templates"
  if test -f "$ftgt"; then
    mv "$ftgt" "$fsrc"
  fi
  qtc -ext .html -dir "$2"
}

function check {
  fsrc="tmp/$1.hashcode"
  ftgt="tmp/$1.hashcode.tmp"

  mkdir -p tmp/

  find "./views" -type f | grep \.html$ | xargs md5sum > "$ftgt"

  if cmp -s "$fsrc" "$ftgt"; then
    if [ "$FORCE" = "force" ]; then
      tmpl $1 $2
    else
      rm "$ftgt"
    fi
  else
    tmpl $1 $2
  fi
}

check "templates" "views"
