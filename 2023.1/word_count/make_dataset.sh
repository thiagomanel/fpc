#!/bin/sh

# it creates a tree dataset
# it creates four directories below the root dir "dataset"
# in each directory, it creates 1000 files. Each file has 154KB

for subDirectoryIndex in $(seq 0 3); do
  subDirectory="dataset/dataset${subDirectoryIndex}"
  mkdir -p $subDirectory
  for fileIndex in $(seq 0 999); do
    cp file2count.data $subDirectory/file2count.data.$fileIndex
  done
done
