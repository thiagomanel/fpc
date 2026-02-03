#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

DATA_DIR="$BASE_DIR/../dataset"
PATTERN=$1

if [ ! -d "$DATA_DIR" ]; then
  echo "Diretório de entrada não encontrado em $DATA_DIR. Execute primeiro, na raiz do lab, o script ./make_dataset.sh."
  exit 1
fi

echo "Executando versão SERIAL..."
time java -cp "$BASE_DIR/bin" DnaSerialMain "$DATA_DIR" "$PATTERN"
