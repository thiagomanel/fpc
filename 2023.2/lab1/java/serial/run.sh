#!/bin/bash

# Verifica se o número de argumentos é válido
if [[ $# -ne 1 ]]; then
  echo "Usage: $0 matrix_size"
  exit 1
fi

# Executa o programa java
MATRIX_SIZE=$1
time java -cp "bin" Find $MATRIX_SIZE
