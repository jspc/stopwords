# stopwords

This project is inspired by the step https://towardsdatascience.com/classify-toxic-online-comments-with-lstm-and-glove-e455a58da9c7#9705, which on Jupyter on my machine is taking about 17 minutes.

This project can achieve the same output in about 30 seconds.

## Installation

```bash
$ go get github.com/jspc/stopwords
```

## Usage

The basic invocation looks like:

```bash
$ stopwords --column data --input testdata/input.csv --output my-output.csv
```

Where flags are:

```bash
$ stopwords -h
Usage of ./stopwords:
  -column string
        Column to clean (default "text")
  -input string
        CSV file to parse (default "in.csv")
  -output string
        CSV file to write to (default "out.csv")
```
