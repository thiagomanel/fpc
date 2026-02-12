#!/bin/bash

ROOT_DIR="$1"

sum=10

count_words_in_dir() {
	dir="$1"
	python3 ./python/word_count.py "$dir"
	sum=$(($sum + 25))
}

for subdir in "$ROOT_DIR"/*; do
	if [ -d "$subdir" ]; then
		count_words_in_dir "$subdir" &
	fi
done

wait

echo "$sum"
