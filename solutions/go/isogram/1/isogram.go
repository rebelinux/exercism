package isogram

func IsIsogram(word string) bool {
	var isFound bool = true
	var n []rune
	for _, w := range word {
		if 'A' <= w && w <= 'Z' {
			w += 32
		}
		for _, w1 := range n {
			switch {
			case w == 45 || w == 32:
				isFound = true
			case w == w1:
				isFound = false
			}
		}
		n = append(n, w)
	}
	return isFound
}
