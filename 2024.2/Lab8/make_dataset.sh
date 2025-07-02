#!/bin/bash

nfiles=10
nlines_per_file=10
nwords_per_line=10
outputdir="dataset"

function rnd_str () {
   rword=$(awk 'BEGIN { srand(); } { words[NR] = $0 } END { print words[int(rand() * NR) + 1] }' /usr/share/dict/words)
   echo $rword
}

mkdir -p $outputdir

for i in `seq 1 $nfiles`
do
    for j in `seq 1 $nlines_per_file`
    do
        line=""
	for j in `seq 1 $nwords_per_line`
        do
            rstr=$(rnd_str)
	    line=$line" "$rstr
        done
        echo $line >> $outputdir/file_$i
    done
done
