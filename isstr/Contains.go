package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/stringindex"
)

// Contains
//
// Results true if the search text contains anywhere in the text.
//
// Returns true
//
//  - if wholeText contains anywhere the search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
//
// For better performance use `...Ptr` version of the method.
func Contains(
	wholeText, containsSearch string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	return stringindex.Of(
		wholeText,
		containsSearch,
		startsAt,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
