package strhelper

import "strings"

// Returns the first index of the findingString in s
// startsAt cannot be negative
// If found returns the index from first, if not then returns -1
func IndexOfPtr(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if s == nil || findingString == nil {
		panic(searchNullPanicMessage)
	}

	if isCaseSensitive && startsAt == 0 {
		return strings.Index(*s, *findingString)
	}

	if isCaseSensitive {
		return IndexOfCaseSensitive(s, findingString, startsAt)
	}

	return IndexOfCaseInsensitive(s, findingString, startsAt)
}
