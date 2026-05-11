package scrabble

func Score(word string) int {
	var s int
	for _, v := range word {
		switch string(v) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T", "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			s += 1
		case "D", "d", "G", "g":
			s += 2
		case "B", "C", "M", "P", "b", "c", "m", "p":
			s += 3
		case "F", "H", "V", "W", "Y", "f", "h", "v", "w", "y":
			s += 4
		case "K", "k":
			s += 5
		case "J", "X", "j", "x":
			s += 8
		case "Q", "Z", "q", "z":
			s += 10
		}
	}
	return s
}
