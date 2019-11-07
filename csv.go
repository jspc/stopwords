package main

import (
	"encoding/csv"
	"os"
)

// Row represents a row in a document, stored as
// columnName = columnValue
type Row map[string]string

func (r Row) ToSlice(headers []string) (out []string) {
	out = make([]string, len(r))

	for i, header := range headers {
		out[i] = r[header]
	}

	return
}

// Document represents a CSV document
type Document struct {
	headers []string
	rows    []Row
}

func LoadCSV(filename string) (d Document, err error) {
	contents, err := os.Open(filename)
	if err != nil {
		return
	}

	defer contents.Close()

	reader := csv.NewReader(contents)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	d.rows = make([]Row, len(records)-1)

	for i, row := range records {
		if i == 0 {
			d.headers = row

			continue
		}

		r := Row{}
		for j, value := range row {
			r[d.headers[j]] = value
		}

		d.rows[i-1] = r
	}

	return
}

func (d Document) WriteCSV(filename string) (err error) {
	outFile, err := os.Create(filename)
	if err != nil {
		return
	}

	records := make([][]string, len(d.rows)+1)

	records[0] = d.headers
	for i, row := range d.rows {
		records[i+1] = row.ToSlice(d.headers)
	}

	w := csv.NewWriter(outFile)
	w.WriteAll(records)

	return w.Error()
}
