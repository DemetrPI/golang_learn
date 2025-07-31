package parsinglogfiles
import (
    "regexp"
    "fmt"
)   
func IsValidLine(text string) bool {
var re = regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
		return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~=*-]+>|<>`).Split(text, -1)
	fmt.Printf("Result, %s", re)
	return re
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
return regexp.MustCompile(`end-of-line[0-9]*`).ReplaceAllString(text, "")
}


func TagWithUserName(lines []string) []string {
	taggedLines := make([]string, 0, len(lines))

	re := regexp.MustCompile(`User\s+(\S+)`)
	for _, line := range lines {
		submatch := re.FindStringSubmatch(line)
		if submatch == nil {
			taggedLines = append(taggedLines, line)
		} else {
			userName := submatch[1]
			taggedLine := fmt.Sprintf("[USR] %s %s", userName, line)
			taggedLines = append(taggedLines, taggedLine)
		}
	}
	return taggedLines
}



