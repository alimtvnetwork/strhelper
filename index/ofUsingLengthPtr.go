package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/indexinternal"
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
func ofUsingLengthPtr(
	s, findingString *string,
	startsAt int,
	wholeTextLength, searchTextLength int,
) int {
	if startsAt == 0 {
		return strings.Index(*s, *findingString)
	}

	return indexinternal.OfCaseSensitiveUsingLength(
		s,
		findingString,
		startsAt,
		wholeTextLength,
		searchTextLength)
}
