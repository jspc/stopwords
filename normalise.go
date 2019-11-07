package main

import (
	"regexp"
	"strings"
)

var (
	re    = regexp.MustCompile(`[^\w\s]`)
	empty = []byte("")
)

func Normalise(s string) string {
	sb := make([]string, 0)

	// Strip new lines
	s = strings.ReplaceAll(s, "\n", " ")

	// Remove non char/ whitespace
	s = string(re.ReplaceAll([]byte(s), empty))

	// Remove errant double spaces
	s = strings.ReplaceAll(s, "  ", " ")

	// Lower case
	s = strings.ToLower(s)

	// Trim whitespace
	s = strings.Trim(s, " ")

	for _, word := range strings.Split(s, " ") {
		if !words.Contains(word) {
			sb = append(sb, word)
		}
	}

	return strings.Join(sb, " ")
}
