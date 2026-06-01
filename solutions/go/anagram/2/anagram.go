package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var anagrams []string
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(s1, s2 string) bool {
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)

	if len(s1) != len(s2) || s1 == s2 {
		return false
	}
	
	counts := make(map[rune]int)
	for _, r := range s1 {
		counts[r]++
	}
	for _, r := range s2 {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}
