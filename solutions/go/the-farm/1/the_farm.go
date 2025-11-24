package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	number  int
	message string
}

var errDeterminingAmount = errors.New("error determining amount of fodder")
var errDeterminingFactor = errors.New("error determining fattening factor")

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.number, e.message)
}

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(f FodderCalculator, cowsNum int) (float64, error) {
	totalAmountOfFodder, err := f.FodderAmount(cowsNum)
	if err != nil {
		return 0, errDeterminingAmount
	}

	fatteningFactor, err := f.FatteningFactor()
	if err != nil {
		return 0, errDeterminingFactor
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
