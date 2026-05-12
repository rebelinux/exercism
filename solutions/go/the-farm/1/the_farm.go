package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
	message string
	details int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	w, err := weightFodder.FodderAmount()

	if w < 0 && err == ErrScaleMalfunction {
		return 0, errors.New("negative fodder")
	}

	if err == ErrScaleMalfunction && w >= 1 {

		return w * 2 / float64(cows), nil

	}
	if err != nil {
		return 0, err
	}
	if w < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{details: cows}
	}

	return w / float64(cows), err

}
