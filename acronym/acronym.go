package acronym

import "strings"

func Abbreviate(s string) string {
	var result strings.Builder
	for i, c := range s {
		if i == 0 || (s[i-1] == ' ' || s[i-1] == '-' || s[i-1] == '_') {
			if c == '_' || c == '-' || c == ' ' {
				continue
			}
			result.WriteString(string(c))

		}
	}
	return strings.ToUpper(result.String())
}
