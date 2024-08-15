#!/bin/bash

args=`find dataset -type f | xargs`

time bash java/serial/run.sh $args
