package thefarm

import (
	"errors"
	"fmt"
)

var errDeterminingAmount = errors.New("amount could not be determined")

type InvalidCowsError struct {
	number  int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.number, e.message)
}

func DivideFood(f FodderCalculator, cowsNum int) (float64, error) {
	totalAmountOfFodder, err := f.FodderAmount(cowsNum)
	if err != nil {
		return 0, errDeterminingAmount
	}

	fatteningFactor, err := f.FatteningFactor()
	if err != nil {
		return 0, errors.New("factor could not be determined")
	} else {
		return totalAmountOfFodder * fatteningFactor / float64(cowsNum), nil
	}
}

func ValidateInputAndDivideFood(f FodderCalculator, cowsNum int) (float64, error) {
	if cowsNum <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(f, cowsNum)
}

func ValidateNumberOfCows(cowsNum int) error {
	if cowsNum < 0 {
		return &InvalidCowsError{
			number:  cowsNum,
			message: "there are no negative cows",
		}
	} else if cowsNum == 0 {
		return &InvalidCowsError{
			number:  cowsNum,
			message: "no cows don't need food",
		}
	}
	return nil
}
