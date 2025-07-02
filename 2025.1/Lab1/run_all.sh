#!/bin/bash

args=`find dataset -type f | xargs`

echo "Executanto o código python serial"
time bash src/python/serial/run.sh $args

echo "Executanto o código python concurrent"
time bash src/python/concurrent/run.sh $args

echo "--------------------------------------------------------------"

echo "Compilando o código java"
bash src/java/serial/build.sh
bash src/java/concurrent/build.sh

echo "Executanto o código Java serial"
time bash src/java/serial/run.sh $args
echo "Executanto o código Java concurrent"
time bash src/java/concurrent/run.sh $args

echo "--------------------------------------------------------------"
echo "Compilando o código C"
bash src/c/serial/build.sh
bash src/c/concurrent/build.sh

echo "Executanto o código C serial"
time bash src/c/serial/run.sh $args
echo "Executanto o código C concurrent"
time bash src/c/concurrent/run.sh $args
