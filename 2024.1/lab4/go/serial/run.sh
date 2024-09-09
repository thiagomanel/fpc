#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# Verifica se pelo menos um argumento foi passado
if [ "$#" -lt 1 ]; then
  echo "Uso: $0 arquivo1 [arquivo2 ...]"
  exit 1
fi

# chama o programa go com os arg passados para o script bash
go run  "$(dirname "$0")/sum.go" "$@"
