package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(ERR|INF|DBG|TRC|WRN|FTL)\]`)

	return re.MatchString(text)

}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<(~*|\**|=*|-*)>|<(-\*~\*-)\>`)

	return re.Split(text, 3)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)^"password"| password"`)

	var c int

	for _, v := range lines {
		if re.FindStringIndex(v) != nil {
			c += 1
		}
	}

	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\w+`)

	for i, v := range lines {
		if re.FindStringSubmatch(v) != nil {
			match := re.FindStringSubmatch(v)

			re := regexp.MustCompile(`User+\s*`)

			lines[i] = fmt.Sprintf("%s %s", re.ReplaceAllString(match[0], "[USR] "), v)
		}
	}

	return lines

}
