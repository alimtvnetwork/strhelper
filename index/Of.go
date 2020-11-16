package index

import (
	"strings"
)

// Returns the first index of the findingString in s
// startsAt cannot be negative
// If found returns the index from first, if not then returns -1
func Of(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive && startsAt == 0 {
		return strings.Index(s, findingString)
	}

	if isCaseSensitive {
		return OfCaseSensitive(&s, &findingString, startsAt)
	}

	return OfCaseInsensitive(&s, &findingString, startsAt)
}
