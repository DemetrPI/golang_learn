package scrabble

import (
	
	"unicode"
)

func Score(word string) int {
	count := 0
	scoresTable := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	var scoreMap = func() map[rune]int {
		m := make(map[rune]int)
		for score, letters := range scoresTable {
			for _, letter := range letters {
				m[letter] = score
			}
		}
		return m
	}()

	for _, letter := range word {
		letter = unicode.ToUpper(letter)
		if letterScore, ok := scoreMap[letter]; ok {
			count += letterScore
		}
	}
	return count
}
