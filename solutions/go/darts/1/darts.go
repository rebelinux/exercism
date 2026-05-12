package darts

import (
	"math"
)

func Score(x, y float64) int {

	s := math.Sqrt(x*x + y*y)

	switch {
	case s > 10:
		return 0
	case s > 5:
		return 1
	case s > 1:
		return 5
	default:
		return 10
	}
}
