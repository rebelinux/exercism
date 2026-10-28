package microblog

func Truncate(phrase string) string {
	if len(phrase) > 5 {
		convert := []rune(phrase)
		return string(convert[:5])
	}
	return phrase
}
