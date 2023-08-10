# Word Count - GO

[word_count.go](./word_count.go) is the base of a Go program to solve the word count problem.

Before running, make sure you have [Go installed](https://go.dev/doc/install) and the dataset is ready (see below).

## Setting up the dataset

To configure the dataset, use the script [make_dataset.sh](../make_dataset.sh), from [2023.1/word_count](../):

```bash
./make_dataset.sh
```

## Running the word count

From [2023.1/word_count](../), you can run the program with:

```bash
go run ./go/word_count.go dataset
```

In this case, `dataset` is the directory where the dataset is located.
