package strhelper

import (
	"strings"
)

// Returns the first index of the findingString in s
//
// Returns Index
//  - If text is found and nothing is invalid like (none is nil)
//
// Returns -1
//  - When not found or invalid case.
//
// Conditions (for panic):
//  - s or search should NOT be nil.
//  - startsAt cannot be negative.
//  - startsAt larger than the content length.
func IndexOfPtr(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if s == nil || findingString == nil {
		panic(searchNullPanicMessage)
	}

	length := len(*s)
	if startsAt < 0 || length-1 < startsAt {
		startAtIndexFailed(startsAt, length)
	}

	if isCaseSensitive && startsAt == 0 {
		return strings.Index(*s, *findingString)
	}

	if isCaseSensitive {
		return IndexOfCaseSensitive(s, findingString, startsAt)
	}

	return IndexOfCaseInsensitive(s, findingString, startsAt)
}
