package scrabble

import "strings"

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

	var scoreMap = make(map[rune]int)
	for score, letters := range scoresTable {
		for _, letter := range letters {
			scoreMap[letter] = score
		}
	}
	word = strings.ToUpper(word)
	for _, letter := range word {
		if letterScore, ok := scoreMap[letter]; ok {
			count += letterScore
		}
	}
	return count
}
