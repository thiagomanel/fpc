#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")


# Verifica se pelo menos um argumento foi passado
if [ "$#" -lt 1 ]; then
  echo "Usage: $0 path1 [path2 ...]"
  exit 1
fi

# Passa todos os argumentos para o script Python
python3 $BASE_DIR/sum.py "$@"
