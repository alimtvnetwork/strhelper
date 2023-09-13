package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/isstrinternal"
)

// StartsWith
//
// Results true for starts with.
//
// Returns true
//
//  - if wholeText starts with search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
func StartsWith(
	wholeText, startsWith string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	searchLength := len(startsWith)
	wholeTextLength := len(wholeText)

	if searchLength == constants.Zero {
		return wholeTextLength == constants.Zero && startsAt == constants.Zero || wholeTextLength-1 >= startsAt
	}

	if wholeTextLength == constants.Zero {
		return searchLength == constants.Zero && startsAt == constants.Zero
	}

	textLength := wholeTextLength - startsAt

	if searchLength > textLength {
		return false
	}

	if isCaseSensitive {
		return isstrinternal.StartsWith(
			wholeText,
			startsWith,
			startsAt,
			wholeTextLength,
			searchLength)
	}

	// insensitive
	return isStartsWithInsensitiveInternal(
		wholeText,
		startsWith,
		startsAt,
		wholeTextLength,
		searchLength)
}
