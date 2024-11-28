#!/bin/bash

nfiles=100
nwords_per_file=10
outputdir="dataset"

function rnd_str () {

    chars="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    random_string=""
    for i in {1..10}; do
        random_string="${random_string}${chars:RANDOM%${#chars}:1}"
    done
   echo "$random_string"
}

for i in `seq 1 $nfiles`;
do
    for j in `seq 1 $nwords_per_file`;
    do
        rstr=$(rnd_str)
        echo $rstr >> $outputdir/file_$i
    done
done
