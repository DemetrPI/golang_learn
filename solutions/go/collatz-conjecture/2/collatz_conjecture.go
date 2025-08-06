package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	err := fmt.Errorf("invalid input: %v, number must be posiive", n)
	if n <= 0 {
		return 0, err
	} else {
		counter := 0
		for {
			switch {
			case n == 1:
				return counter, nil
			case n%2 == 0:
				n = n / 2
				counter++
			default:
				n = n*3 + 1
				counter++
			}
		}
	}
}
