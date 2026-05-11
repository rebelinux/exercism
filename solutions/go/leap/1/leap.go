package leap

// IsLeapYear calculate the leap year using the following fules:
// on every year that is evenly divisible by 4
// except every year that is evenly divisible by 100
// unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {
	var y bool

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				y = true
			}
		} else {
			y = true
		}
	} else {
		y = false
	}

	return y
}
