package hamming

import (
	"fmt"
	"strings"
)

type customErr error

func Distance(a, b string) (int, error) {
	var s1 []rune
	var diference []int
	var error customErr

	if len(a) == len(b) {
		for _, v := range a {
			s1 = append(s1, v)
		}

		for i, v := range b {
			if strings.Compare(string(v), string(s1[i])) == 1 || strings.Compare(string(v), string(s1[i])) == -1 {
				diference = append(diference, 1)

			} else {
				error = nil
			}
		}
	} else {
		error = fmt.Errorf("Error")
	}

	return len(diference), error
}
