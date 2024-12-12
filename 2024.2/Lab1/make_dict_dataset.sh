#!/bin/bash

nfiles=100
nwords_per_file=10
outputdir="dataset"

function rnd_str () {
   rword=$(awk 'BEGIN { srand(); } { words[NR] = $0 } END { print words[int(rand() * NR) + 1] }' /usr/share/dict/words)
   echo $rword
}

for i in `seq 1 $nfiles`
do
    for j in `seq 1 $nwords_per_file`
    do
        rstr=$(rnd_str)
        echo $rstr-$j >> $outputdir/file_$i
    done
done
