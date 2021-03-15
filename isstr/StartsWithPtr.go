package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// Results true for starts with.
//
// Returns true
//
//  - if wholeText starts with search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
func StartsWithPtr(
	wholeText, startsWith *string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	if wholeText == nil || startsWith == nil {
		panic("Either wholeText or startsWith is nil. Please provide valid string at least EmptyString(\"\")")
	}

	searchLength := len(*startsWith)
	wholeTextLength := len(*wholeText)

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
		return isstrinternal.IsStartsWith(
			wholeText,
			startsWith,
			startsAt)
	}

	// insensitive
	return isStartsWithInsensitiveInternal(
		wholeText,
		startsWith,
		startsAt)
}
