#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# Verifica se pelo menos um argumento foi passado
if [ "$#" -lt 1 ]; then
  echo "Uso: $0 arquivo1 [arquivo2 ...]"
  exit 1
fi

# chama o programa java com os arg passados para o script bash
java -cp $BASE_DIR/bin/ Sum "$@"
