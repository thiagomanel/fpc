#!/bin/bash

BASE_DIR=$(dirname -- "$(readlink -f -- "$0")")

if [ "$#" -ne 5 ]; then
    echo "Uso: $0 <num_producers> <max_items_per_producer> <producing_time> <num_consumers> <consuming_time>"
    exit 1
fi

PRODUCERS=$1
MAX_ITEMS_PRODUCERS=$2
PRODUCING_TIME=$3
CONSUMERS=$4
CONSUMING_TIME=$5

java -cp "$BASE_DIR/java/bin" Main "$PRODUCERS" "$MAX_ITEMS_PRODUCERS" "$PRODUCING_TIME" "$CONSUMERS" "$CONSUMING_TIME"
