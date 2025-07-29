package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	var re = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~=*-]+>|<>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		re := regexp.MustCompile(`"[^"]*\b(?i)password\b[^"]*"`)
		match := re.MatchString(line)
		if match {
			count++
		}
	}
	return count
}
func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, 0, len(lines))
	re := regexp.MustCompile(`User\s+(\S+)`)
	for _, line := range lines {
		if submatch := re.FindStringSubmatch(line); submatch != nil {
			userName := submatch[1]
			tagged := fmt.Sprintf("[USR] %s %s", userName, line)
			taggedLines = append(taggedLines, tagged)
		} else {
			taggedLines = append(taggedLines, line)
		}
	}
	return taggedLines
}
