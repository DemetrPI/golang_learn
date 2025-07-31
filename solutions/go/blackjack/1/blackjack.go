package blackjack

// ParseCard converts a card string to its corresponding integer value.
func ParseCard(card string) int {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case playerScore == 22:
		return "P"
	case playerScore == 21:
		if dealerCardValue != 10 && dealerCardValue != 11 {
			return "W"
		}
		return "S"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore >= 12 && playerScore <= 16:
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	default: // playerScore <= 11
		return "H"
	}
}
