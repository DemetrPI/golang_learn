package scrabble

import "testing"

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type ScrabbleTest struct {
	description string
	input       string
	expected    int
}

var ScrabbleScoreTests = []ScrabbleTest{
	{
		description: "lowercase letter",
		input:       "a",
		expected:    1,
	},
	{
		description: "uppercase letter",
		input:       "A",
		expected:    1,
	},
	{
		description: "valuable letter",
		input:       "f",
		expected:    4,
	},
	{
		description: "short word",
		input:       "at",
		expected:    2,
	},
	{
		description: "short, valuable word",
		input:       "zoo",
		expected:    12,
	},
	{
		description: "medium word",
		input:       "street",
		expected:    6,
	},
	{
		description: "medium, valuable word",
		input:       "quirky",
		expected:    22,
	},
	{
		description: "long, mixed-case word",
		input:       "OxyphenButazone",
		expected:    41,
	},
	{
		description: "english-like word",
		input:       "pinata",
		expected:    8,
	},
	{
		description: "empty input",
		input:       "",
		expected:    0,
	},
	{
		description: "entire alphabet available",
		input:       "abcdefghijklmnopqrstuvwxyz",
		expected:    87,
	},
}

func TestScore(t *testing.T) {
	for _, tc := range ScrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Score(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}
func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range ScrabbleScoreTests {
			Score(test.input)
		}
	}
}