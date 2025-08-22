package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) !=len(b) {
		return -1, errors.New("incorrect string lenght")
	} else {
			for i := range a {
			if a[i] !=b[i] {
				count++
			} 		
	}
	return count, nil
	}
}
