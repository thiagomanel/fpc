#!/bin/bash

if [ $# -lt 2 ]; then
  echo "Erro: Informe o caminho da imagem e o número de threads como parâmetros."
  exit 1
fi

IMAGE_PATH=$1
THREADS=$2

javac ImageMeanFilter.java
time java ImageMeanFilter "$IMAGE_PATH" "$THREADS"
