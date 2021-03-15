package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
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
	if searchTextLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	for i := startsAt; i < wholeTextLength; i++ {
		if wholeTextLength-i < searchTextLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.IsStartsWith(s, findingString, i) {
			return i
		}
	}

	return constants.InvalidNotFoundCase

}
