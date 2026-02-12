#!/bin/bash

ROOT_DIR="$1"

count_words_in_dir() {
	dir="$1"
	go run ./go/word_count.go "$dir"
}

for subdir in "$ROOT_DIR"/*; do
	if [ -d "$subdir" ]; then
		count_words_in_dir "$subdir" &
	fi
done

wait
