// package leap

// // IsLeapYear should have a comment documenting it.
// func IsLeapYear(year int) bool {
// 	switch {
// 	case year%4 != 0:
// 		return false
// 	case year%400 == 0:
// 		return true
// 	case year%4 == 0 && year%100 == 0:
// 		return false
// 	case year%100 == 0 && year%400 == 0:
// 		return true
// 	default:
// 		return true
// 	}
// }

package leap

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}