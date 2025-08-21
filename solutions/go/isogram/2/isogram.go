package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, letter := range strings.ToLower(word) {
		if letter == '-' || letter == ' ' {
			continue
		}
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}
