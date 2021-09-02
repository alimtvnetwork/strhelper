package isstr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/strhelper/index"
)

// Results true if the search text contains anywhere in the text.
//
// Returns true
//
//  - if wholeText contains any where the search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
func ContainsPtr(
	wholeText, containsSearch *string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	return index.OfPtr(
		wholeText,
		containsSearch,
		startsAt,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
