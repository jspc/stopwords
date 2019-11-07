package main

import (
	"fmt"
	"testing"
)

func TestNormalise(t *testing.T) {
	for idx, test := range []struct {
		input  string
		expect string
	}{
		{"Hello, World!", "world"},
		{"This doesn't always work; for instance: Among My Swan is one of my favourite albums, but I bet we remove the word 'Among'", "doesnt work instance swan favourite albums i bet remove word"},
	} {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := Normalise(test.input)
			if test.expect != got {
				t.Errorf("expected %q, received %q", test.expect, got)
			}
		})
	}
}
