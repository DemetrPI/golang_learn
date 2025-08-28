package luhn

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func NumbersAndStrings(id string) (bool, error) {
	return regexp.MatchString(`^[0-9 ]*$`, id)
}

func Valid(id string) bool {
	num := regexp.MustCompile(`\D+`)
	if len(id) <= 1 {
		return false
	}
    if id == " 0"{
		return false
	}
	isValid, _ := NumbersAndStrings(id)
	if !isValid {
		return false
	}

	id = num.ReplaceAllString(id, "")
	id = Reverse(id)
	result := 0
	for i, v := range id {
		digit, _ := strconv.Atoi(string(v))
		if i%2 != 0 {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		result = result + digit
	}
	return result%10 == 0
}
