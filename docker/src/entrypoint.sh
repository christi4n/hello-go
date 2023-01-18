#!/usr/bin/env bash

# set -e instructs bash to immediately exit 
# if any command [1] has a non-zero exit status
# set -o  prevents errors in a pipeline from being masked.
# src: https://gist.github.com/mohanpedala/1e2ff5661761d3abd0385e8223e16425
set -eo pipefail
echo ">>> Running command"
echo ""
bash -c "set -e;  set -o pipefail; $1"