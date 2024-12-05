#!/bin/bash

if [ $# -lt 1 ]; then
  echo "Erro: Informe o caminho da imagem como parâmetro."
  exit 1
fi

IMAGE_PATH=$1

javac ImageMeanFilter.java
time java ImageMeanFilter "$IMAGE_PATH"
