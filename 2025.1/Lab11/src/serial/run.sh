#!/bin/bash

BASE_DIR=$(dirname -- "$(readlink -f -- "$0")")

if [ "$#" -ne 4 ]; then
    echo "Uso: $0 <num_producers> <producing_time> <num_consumers> <consuming_time>"
    exit 1
fi

PRODUCERS=$1
PRODUCING_TIME=$2
CONSUMERS=$3
CONSUMING_TIME=$4

java -cp "$BASE_DIR/java/bin" Main "$PRODUCERS" "$PRODUCING_TIME" "$CONSUMERS" "$CONSUMING_TIME"
