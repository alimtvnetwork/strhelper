package strhelper

import (
	"strings"
)

// Returns the first index of the findingString in s
// startsAt cannot be negative
// If found returns the index from first, if not then returns -1
func IndexOf(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive && startsAt == 0 {
		return strings.Index(s, findingString)
	}

	if isCaseSensitive {
		return IndexOfCaseSensitive(&s, &findingString, startsAt)
	}

	return IndexOfCaseInsensitive(&s, &findingString, startsAt)
}
