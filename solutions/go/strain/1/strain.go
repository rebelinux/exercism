package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {

	var cp Ints

	for _, v := range i {
		if filter(v) {
			cp = append(cp, v)
		}
	}

	return cp
}

func (i Ints) Discard(filter func(int) bool) Ints {

	var cp Ints

	for _, v := range i {
		if !filter(v) {
			cp = append(cp, v)
		}
	}

	return cp
}

func (l Lists) Keep(filter func([]int) bool) Lists {

	var cp Lists

	for _, v := range l {

		if filter(v) {
			cp = append(cp, v)
		}
	}

	return cp
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var cp Strings

	for _, v := range s {

		if filter(v) {
			cp = append(cp, v)
		}
	}

	return cp
}
