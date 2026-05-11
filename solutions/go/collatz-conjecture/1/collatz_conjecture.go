package collatzconjecture

import "errors"

func CollatzConjecture(n int) (steps int, err error) {

	if n < 1 {
		return -1, errors.New("zero and negative numbers are unsupported")
	}
	for n != 1 {
		steps++
		if n&1 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return
}
