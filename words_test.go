package main

import (
	"testing"
)

func TestWords_Contains(t *testing.T) {
	t.Run("word is contained in stopwords", func(t *testing.T) {
		if !words.Contains("indicated") {
			t.Errorf("expected true, received false")
		}
	})

	t.Run("word is not a stopword", func(t *testing.T) {
		if words.Contains("balloon") {
			t.Errorf("expected false, received true")
		}
	})
}
