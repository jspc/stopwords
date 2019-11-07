package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestLoadCSV(t *testing.T) {
	expect := Document{
		headers: []string{"id", "score", "data"},
		rows:    []Row{{"id": "1", "score": "11", "data": "This is a message with some stuff in it!"}},
	}

	for _, test := range []struct {
		name        string
		filename    string
		expect      Document
		expectError bool
	}{
		{"Happy path", "testdata/input.csv", expect, false},
		{"Missing file", "testdata/nonsuch.csv", Document{}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := LoadCSV(test.filename)

			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if !reflect.DeepEqual(test.expect, got) {
				t.Errorf("expected %+v, got %+v", test.expect, got)
			}
		})
	}
}

func TestSaveCSV(t *testing.T) {
	d := Document{
		headers: []string{"id", "score", "data"},
		rows:    []Row{{"id": "1", "score": "11", "data": "This is a message with some stuff in it!"}},
	}

	output, err := ioutil.TempFile("", "csv_test")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	err = d.WriteCSV(output.Name())
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
}
