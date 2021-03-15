package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// Results true for ends with search text.
//
// Returns true
//
//  - if wholeText starts from the last with search text comparison.
//  - if lastIndexIncreasedBy mentioned then last len(wholeText)-lastIndexIncreasedBy
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - lastIndexIncreasedBy cannot be negative
//
// lastIndexIncreasedBy:
//  - `2` represents len(wholeText)-2
//  - `0` represents start comparison from the end for both of the text.
func EndsWithPtr(
	wholeText, search *string,
	lastIndexIncreasedBy int,
	isCaseSensitive bool,
) bool {
	if wholeText == nil || search == nil {
		panic("Either wholeText or search is nil. Please provide valid string at least EmptyString(\"\")")
	}

	searchLength := len(*search)
	wholeTextLength := len(*wholeText)

	if searchLength == constants.Zero {
		return (wholeTextLength == constants.Zero && lastIndexIncreasedBy == constants.Zero) ||
			wholeTextLength-1 >= lastIndexIncreasedBy
	}

	if wholeTextLength == constants.Zero {
		return searchLength == constants.Zero && lastIndexIncreasedBy == constants.Zero
	}

	textLength := wholeTextLength - lastIndexIncreasedBy
	if searchLength > textLength {
		return false
	}

	if isCaseSensitive {
		return isstrinternal.IsEndsWithInternal(
			wholeText,
			search,
			lastIndexIncreasedBy)
	}

	// insensitive
	return isEndsWithInsensitiveInternal(
		wholeText,
		search,
		lastIndexIncreasedBy)
}
