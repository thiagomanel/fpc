#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 path1 [path2 ...]"
  exit 1
fi

# chama o programa c com os arg passados para o script bash
"$BASE_DIR"/sum "$@"