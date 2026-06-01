package microblog

func Truncate(phrase string) string {
	runes := []rune(phrase)
	char := make([]rune, len(runes))
	copy(char, runes)
	if len(char) > 5 {
		return string(char[:5])
	}
	return string(char)
}

