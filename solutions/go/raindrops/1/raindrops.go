package raindrops

import (
	"fmt"
)

func Convert(number int) string {

	var x string

	if number%3 == 0 {
		x += "Pling"
	}
	if number%5 == 0 {
		x += "Plang"
	}
	if number%7 == 0 {
		x += "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		x += fmt.Sprint(number)
	}

	return fmt.Sprint(fmt.Sprint(x))
}
