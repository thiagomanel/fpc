#!/bin/bash

args=`find data -type f | xargs`

echo "Compilando o código java"
bash src/serial/build.sh
bash src/concurrent/build.sh

echo "Executanto o código Java serial"
time bash src/serial/run.sh $args
echo "Executanto o código Java concurrent"
time bash src/concurrent/run.sh $args
