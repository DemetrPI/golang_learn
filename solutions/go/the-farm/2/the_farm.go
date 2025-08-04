package thefarm

import (
	"fmt"
"errors"
)
var errInvalidNumberOfCows = errors.New("invalid number of cows")


type InvalidCowsError struct {
	number int
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
		return 0, errDeterminingFactor
	} 
		return totalAmountOfFodder * fatteningFactor / float64(cowsNum), nil
	
}

func ValidateInputAndDivideFood(f FodderCalculator, cowsNum int) (float64, error) {
	if err:= ValidateNumberOfCows(cowsNum); err!=nil {
		return 0, errInvalidNumberOfCows
	}
	return DivideFood(f, cowsNum)
}

func ValidateNumberOfCows(cowsNum int) error {
	switch {
	case cowsNum < 0:
		{
			return &InvalidCowsError{
				number:  cowsNum,
				message: "there are no negative cows",
			}
		}
	case cowsNum == 0:
		return &InvalidCowsError{
			number:  cowsNum,
			message: "no cows don't need food",
		}
	default:
		return nil
	}
}
