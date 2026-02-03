#!/bin/bash

if [ $# -lt 1 ]; then
  echo "Erro: Informe o caminho da imagem como parâmetro."
  exit 1
fi

IMAGE_PATH=$1

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

mkdir -p "$BASE_DIR/bin"

echo "Compilando código..."
javac -d "$BASE_DIR/bin" "$BASE_DIR"/*.java

echo "Executando..."
time java -cp "$BASE_DIR/bin" ImageMeanFilter "$IMAGE_PATH"
