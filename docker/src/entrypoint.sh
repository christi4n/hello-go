#!/usr/bin/env bash

set -eo pipefail
echo ">>> Running command"
echo ""
bash -c "set -e;  set -o pipefail; $1"