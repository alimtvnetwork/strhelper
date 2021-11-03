package stringindex

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
// Recommendation:
//  - Use ptr version for performance.
func Of(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive && startsAt == 0 {
		return strings.Index(s, findingString)
	}

	if isCaseSensitive {
		return OfCaseSensitive(
			s,
			findingString,
			startsAt)
	}

	return OfCaseInsensitive(
		s,
		findingString,
		startsAt)
}
