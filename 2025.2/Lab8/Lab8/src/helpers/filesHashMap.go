package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

func sum(filePath string) (int, error) {
	data, err := readFile(filePath)
	if err != nil {
		return 0, err
	}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	return _sum, nil
}

func listarArquivos(diretorio string) []string {
	arquivos, err := ioutil.ReadDir(diretorio)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, arquivo := range arquivos {
		if !arquivo.IsDir() {
			result = append(result, arquivo.Name())
		}
	}
	return result
}

// Versão SEM concorrência: sem goroutines e sem channels.
// Mantém a funcionalidade de gerar o "hash map" (na prática, checksums).
func generateFilesHashMap() map[string][]int {
	diretorio := "../dataset/"

	if _, err := os.Stat(diretorio); os.IsNotExist(err) {
		log.Fatalf("O diretório %s não existe", diretorio)
	}

	files := listarArquivos(diretorio)

	hashs := make(map[string][]int)
	for _, name := range files {
		filePath := diretorio + name
		fileSum, err := sum(filePath)
		if err != nil {
			fileSum = 0
		}
		hashs[filePath] = append(hashs[filePath], fileSum)
	}

	return hashs
}
