package logs

import (
	"strings"
	"unicode/utf8"
)

var charApp = map[rune]string{
  '❗': "recommendation",
  '🔍': "search",
  '☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}
// Another solution
func ApplicationLog(log string) string {
  for _, char := range log {
    // Check if the char exists as a key in charApp
    // and if it found then return the application
    // string for it
    if app, found := charApp[char]; found {
      return app
    }
  }
 
  return "default"
}

// Relace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
