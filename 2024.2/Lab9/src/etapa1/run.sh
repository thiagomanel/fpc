#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

java -cp $BASE_DIR/bin/ Main
