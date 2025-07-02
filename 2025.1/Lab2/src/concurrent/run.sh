#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# Verifica se pelo menos um argumento foi passado
if [ "$#" -lt 1 ]; then
  echo "Uso: $0 arquivo1 [arquivo2 ...]"
  exit 1
fi

## Compila o código direcionando a saída para a pasta bin
#javac -d $BASE_DIR/bin $BASE_DIR/Sum.java

# chama o programa java com os arg passados para o script bash
time java -cp $BASE_DIR/bin/ Deconvolution "$@"
