#!/usr/bin/env bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

OUT_DIR="$BASE_DIR/dataset"
NUM_FILES=$1
NUM_SEQS_PER_FILE=100
MIN_LEN=80
MAX_LEN=120

mkdir -p "$OUT_DIR"

letters=(A C G T)

echo "Gerando $NUM_FILES arquivos em '$OUT_DIR'..."

for ((f = 0; f < NUM_FILES; f++)); do
  FILE="$OUT_DIR/dna_$f.txt"
  > "$FILE"
  for ((i = 0; i < NUM_SEQS_PER_FILE; i++)); do
    len=$((RANDOM % (MAX_LEN - MIN_LEN + 1) + MIN_LEN))
    seq=""
    for ((j = 0; j < len; j++)); do
      seq+="${letters[$((RANDOM % 4))]}"
    done
    echo "$seq" >> "$FILE"
  done
  echo "Arquivo gerado: $FILE"
done

echo "Concluído."
