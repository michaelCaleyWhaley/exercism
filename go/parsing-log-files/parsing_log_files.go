package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	pos := re.FindIndex([]byte(text))

	if len(pos) != 0 {
		return pos[0] == 0
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<[\*~=\-]*>)`)
	length := len(re.FindIndex([]byte(text))) + 1
	splitString := re.Split(text, length)
	return splitString
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*(password).*"`)

	count := 0
	for _, line := range lines {
		length := len(re.FindIndex([]byte(line)))
		if length > 0 {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line[0-9]*)`)
	return string(re.ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User[[:space:]]+[a-zA-Z0-9]+)`)
	userNameRe := regexp.MustCompile(`(?i)([a-z]+[0-9]+)`)

	for index, line := range lines {
		match := re.Find([]byte(line))

		if len(match) != 0 {
			username := userNameRe.Find(match)
			lines[index] = "[USR] " + string(username) + " " + line
		}

	}

	return lines
}
