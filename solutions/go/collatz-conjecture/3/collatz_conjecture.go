package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid input: %v, number must be posiive", n)
	} 
    counter := 0
		for n!=1 {
			if n%2 == 0{
				n = n / 2
                } else {
				n = n*3 + 1                
                }
				counter++
			}
				return counter, nil
		}

