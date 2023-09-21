package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Count the number of words in `fileContent`.
func wc(fileContent string) int {
	words := strings.Fields(fileContent)
	return len(words)
}

// Count the number of words in the file at `filePath`.
func wc_file(filePath string) int {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return wc(string(fileContent))
}

// Count the number of words in all files directly within `directoryPath`.
// Files in subdirectories are not considered.
func wc_dir(directoryPath string) int {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	numberOfWords := 0

	for _, file := range files {
		if !file.IsDir() {
			filePath := directoryPath + "/" + file.Name()
			numberOfWords += wc_file(filePath)
		}
	}

	return numberOfWords
}

// Calculate the number of words in the files stored under the directory name
// available at argv[1].
//
// Assume a depth 3 hierarchy:
//   - Level 1: root
//   - Level 2: subdirectories
//   - Level 3: files
//
// root
// ├── subdir 1
// │     ├── file
// │     ├── ...
// │     └── file
// ├── subdir 2
// │     ├── file
// │     ├── ...
// │     └── file
// ├── ...
// └── subdir N
// │     ├── file
// │     ├── ...
// │     └── file
func main() {
	rootPath := os.Args[1]

	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	numberOfWords := 0

	for _, file := range files {
		if file.IsDir() {
			directoryPath := rootPath + "/" + file.Name()
			numberOfWords += wc_dir(directoryPath)
		}
	}

	fmt.Println(numberOfWords)
}
