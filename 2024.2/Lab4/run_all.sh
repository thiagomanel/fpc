#!/bin/bash

args=`find dataset -type f | xargs`

echo "Building Java serial implementation"
bash src/serial/java/build.sh

echo "Running Serial implementation"
time bash src/serial/java/run.sh $args

echo "Building Java Concurrent implementation"
bash src/concurrent/java/build.sh

echo "Running Concurrent implementation"
time bash src/concurrent/java/run.sh $args
