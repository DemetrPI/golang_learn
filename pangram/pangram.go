package pangram

import "strings"

func IsPangram(s string) bool {
	toLowerCase := strings.ToLower(s)
	seen := make(map[rune]bool)
	for _, r := range toLowerCase {
		if r >= 'a' && r <= 'z' {
			seen[r] = true
		}
	}
	return len(seen) == 26
}