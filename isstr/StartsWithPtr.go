package isstr

import (
	"gitlab.com/evatix-go/strhelper/internal/isinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
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

	if searchLength == strconst.Zero {
		return wholeTextLength == strconst.Zero && startsAt == strconst.Zero || wholeTextLength-1 >= startsAt
	}

	if wholeTextLength == strconst.Zero {
		return searchLength == strconst.Zero && startsAt == strconst.Zero
	}

	textLength := wholeTextLength - startsAt

	if searchLength > textLength {
		return false
	}

	if isCaseSensitive {
		return isinternal.IsStartsWithInternal(
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
