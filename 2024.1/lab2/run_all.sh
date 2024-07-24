#!/bin/bash

args=`find dataset -type f | xargs`

time bash python/serial/run.sh $args
time bash python/concurrent/run.sh $args

time bash java/serial/run.sh $args
time bash java/concurrent/run.sh $args