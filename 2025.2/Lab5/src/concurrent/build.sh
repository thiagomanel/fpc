#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

mkdir -p $BASE_DIR/java/bin

javac -d $BASE_DIR/java/bin $BASE_DIR/java/*.java
