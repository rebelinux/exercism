package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	var s1 []rune
	var diference []int

	if len(a) == len(b) {
		for _, v := range a {
			s1 = append(s1, v)
		}

		for i, v := range b {
			if strings.Compare(string(v), string(s1[i])) == 1 || strings.Compare(string(v), string(s1[i])) == -1 {
				diference = append(diference, 1)

			} else {
				return 0, nil
			}
		}
	} else {
		return 0, errors.New("Error in string length")
	}

	return len(diference), nil
}
