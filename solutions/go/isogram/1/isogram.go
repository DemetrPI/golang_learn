package isogram

import "strings"

func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	counts := make(map[rune]int)

	for _, letter := range word {
		if letter == '-' || letter == ' ' {
			continue
		} else {
			counts[letter]++
		}
		if counts[letter] > 1 {
			return false
		}
	}
	return true
}
