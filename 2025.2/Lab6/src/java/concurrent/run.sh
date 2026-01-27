#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# Verifica se pelo menos um argumento foi passado
if [ "$#" -lt 1 ]; then
  echo "Uso: $0 number_of_users"
  exit 1
fi

# chama o programa java com os arg passados para o script bash
time java -cp $BASE_DIR/bin/ WebStatsMain "$@"
