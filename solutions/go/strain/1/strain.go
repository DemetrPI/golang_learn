package strain

func Index[T comparable](element T, list []T) int {
	for i, item := range list {
		if item == element {
			return i
		}
	}
	return -1
}
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
