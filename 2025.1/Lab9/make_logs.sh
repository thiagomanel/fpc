#!/bin/bash

# Uso: ./make_logs.sh NUM_ARQUIVOS TAMANHO
# Exemplo: ./make_logs.sh 5 1000
# -> gera 5 arquivos de log, cada um com 1000 linhas

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

NUM_FILES=$1
SIZE=$2

mkdir -p $BASE_DIR/dataset

if [ -z "$NUM_FILES" ] || [ -z "$SIZE" ]; then
  echo "Uso: $0 NUM_ARQUIVOS TAMANHO"
  exit 1
fi

for i in $(seq 1 $NUM_FILES); do
  FILE="${BASE_DIR}/dataset/log${i}.txt"
  echo "Gerando $FILE com $SIZE linhas..."
  > $FILE
  for j in $(seq 1 $SIZE); do
    METHOD=$((RANDOM % 2))
    if [ $METHOD -eq 0 ]; then
      HTTP_METHOD="GET"
      URL="/index.html"
    else
      HTTP_METHOD="POST"
      URL="/api/payments"
    fi

    STATUS=$((RANDOM % 10))
    if [ $STATUS -lt 8 ]; then
      CODE=200
    else
      CODE=500
    fi

    echo "$HTTP_METHOD $URL $CODE" >> $FILE
  done
done

echo "Concluído!"

