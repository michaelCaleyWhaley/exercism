package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeSlice := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}
	applicationType := "default"
	runePos := -1

	for runeKey, runeValue := range runeSlice {

		indexOfRune := strings.IndexRune(log, runeKey)

		if indexOfRune != -1 && (indexOfRune < runePos || runePos == -1) {
			applicationType = runeValue
			runePos = indexOfRune
		}
	}
	return applicationType
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeSlice := []rune(log)
	for index, value := range runeSlice {
		if value == oldRune {
			runeSlice[index] = newRune
		}
	}
	return string(runeSlice)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
