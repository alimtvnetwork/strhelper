package strhelper

import (
	"strings"
)

// Returns the last index of the findingString in s
// it returns the index where the word starts from not the end of index
// startsAt cannot be negative
// If found returns the index from last, if not then returns -1
func LastIndexOfPtr(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if s == nil || findingString == nil {
		panic(searchNullPanicMessage)
	}

	if isCaseSensitive && startsAt == 0 {
		return strings.LastIndex(*s, *findingString)
	}

	if isCaseSensitive {
		return LastIndexOfCaseSensitive(s, findingString, startsAt)
	}

	return LastIndexOfCaseInsensitive(s, findingString, startsAt)
}
