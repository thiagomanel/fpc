#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

# chama o programa java DistriutedSystem
java -cp $BASE_DIR/bin/ ScenarioBase
