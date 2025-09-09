#!/bin/bash

args=`find dataset -type f | xargs`

echo "Compilando o código java"
bash src/java/serial/build.sh
#bash src/java/concurrent/build.sh

echo "Executanto o código Java serial"
time bash src/java/serial/run.sh $args
#echo "Executanto o código Java concurrent"
#time bash src/java/concurrent/run.sh $args
