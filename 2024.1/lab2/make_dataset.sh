#/bin/bash

# generate $nfiles random content files in a dataset directory
nfiles=$1

mkdir -p dataset

for i in `seq 1 $nfiles`
do
	# Generate a random size between 50k and 72k
    file_size=$((1000000 + RANDOM % 22001))
	head -c $file_size < /dev/urandom > dataset/file.$i
done
