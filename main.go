package main

import (
	"context"
	"flag"
	"log"
	"runtime"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0) - 1
	pool       = semaphore.NewWeighted(int64(maxWorkers))

	input  = flag.String("input", "in.csv", "CSV file to parse")
	output = flag.String("output", "out.csv", "CSV file to write to")
	column = flag.String("column", "text", "Column to clean")
)

func main() {
	flag.Parse()

	inDoc, err := LoadCSV(*input)
	if err != nil {
		log.Panic(err)
	}

	outDoc := Document{
		headers: inDoc.headers,
		rows:    make([]Row, len(inDoc.rows)),
	}

	ctx := context.TODO()
	for idx, row := range inDoc.rows {
		pool.Acquire(ctx, 1)

		go func(i int, r Row) {
			defer func() {
				err := recover()
				if err != nil {
					log.Printf("idx: %d, row: %+v", i, r)

					log.Panic(err)
				}
			}()

			r[*column] = Normalise(r[*column])
			outDoc.rows[i] = r

			pool.Release(1)
		}(idx, row)
	}

	err = outDoc.WriteCSV(*output)
	if err != nil {
		log.Panic(err)
	}
}
