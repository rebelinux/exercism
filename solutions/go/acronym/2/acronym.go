// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var m []byte

	replace := strings.NewReplacer("-", " ", "_", "")

	n := strings.Split(strings.ToUpper(replace.Replace(s)), " ")

	for _, p := range n {
		if p != "" {
			m = append(m, (p[0]))
		}
	}

	return string(m)
}
