package sorting
import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f",f)
}
type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch v := fnb.(type) {
	case FancyNumber:
		var i, _ = strconv.Atoi(v.Value())
		return i
	default:
		return 0
	}
}

// Another possible implementation 
// func ExtractFancyNumber(fnb FancyNumberBox) int {
// 	num, ok := fnb.(FancyNumber)
// 	if ok {
// 		num, _ := strconv.Atoi(num.Value())
// 		return num
// 	}
// 	return 0
// }


// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i := i.(type) {
	case float64:
		return DescribeNumber(i)
	case int:
		return DescribeNumber(float64(i))
	case NumberBox:
		return DescribeNumberBox(i)
	case FancyNumber:
		return DescribeFancyNumberBox(i)
	case FancyNumberBox:
		return DescribeFancyNumberBox(i)
	default:
		return "Return to sender"
	}
}
