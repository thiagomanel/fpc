# Word Count - Java

[WordCount.java](./src/main/java/WordCount.java) is the base of a Java program to solve the word count problem.

Before running, make sure you have [Java 11+](https://www.oracle.com/br/java/technologies/javase/jdk11-archive-downloads.html) and [Gradle installed](https://gradle.org/install/). Furthermore, confirm that the dataset is ready (see below).

## If you don't have gradle installed

```bash
curl -s "https://get.sdkman.io" | bash
```

Open a new terminal, then run this:
```bash
sdk install gradle 8.3
```

## Setting up the dataset

To configure the dataset, use the script [make_dataset.sh](../make_dataset.sh), from [2023.1/word_count](../):

```bash
./make_dataset.sh
```

## Running the word count

From [2023.1/word_count/java](/), you can run the program with:

```bash
gradle runApp
```
