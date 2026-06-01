package strain

func Keep[T any](numbers []T, predicate func(T) bool) []T {
	var result []T
	for _, number := range numbers {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return result
}

func Discard[T any](numbers []T, predicate func(T) bool) []T {
	var result []T
	for _, number := range numbers {
		if !predicate(number) {
			result = append(result, number)
		}
	}
	return result
}
