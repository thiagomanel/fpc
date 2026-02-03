#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# chama o programa java com os arg passados para o script bash
time java -cp $BASE_DIR/bin/ SimpleSerialSolution

