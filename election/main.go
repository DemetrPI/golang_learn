package electionday

import (
	"fmt"
	"strings"
)

// not required for exercism test, as they use this struct in test.
type ElectionResult struct {
	Name  string
	Votes int
}

var counter *int
var result *ElectionResult

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	counter = &initialVotes
	return counter
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter

}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(Name string, Votes int) *ElectionResult {
	result = &ElectionResult{Name, Votes}
	return result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	s := strings.Join([]string{result.Name, "(", fmt.Sprint(result.Votes), ")"}, "")
	return s
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
results[candidate]--
}
