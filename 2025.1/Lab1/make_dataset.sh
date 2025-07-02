#/bin/bash

# generate $nfiles random content files in a dataset directory
nfiles=$1

mkdir -p dataset

for i in `seq 1 $nfiles`
do
	#from 0 to 32K
	file_size=$RANDOM
	head -c $file_size < /dev/urandom > dataset/file.$i
done
