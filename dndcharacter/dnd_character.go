package dndcharacter

import (
	"math"
	"math/rand/v2"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int((math.Floor((float64(score) - 10) / 2)))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	smallest := 6
	result := 0

	for i := 0; i < 4; i++ {
		dice := rand.IntN(6) + 1
		result += dice
		if dice <= smallest {
			smallest = dice
		}
	}
	return result - smallest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {

	char := Character{
		Charisma:     Ability(),
		Constitution: Ability(),
		Dexterity:    Ability(),
		Intelligence: Ability(),
		Strength:     Ability(),
		Wisdom:       Ability(),
	}
	char.Hitpoints = (10 + Modifier(char.Constitution))
	return char
}
