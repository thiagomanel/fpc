#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# Compila o código direcionando a saída para mesmo local de sum.c 
gcc -o $BASE_DIR/sum $BASE_DIR/sum.c