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
//  - if contentLengthDecreasedBy mentioned then last len(wholeText)-contentLengthDecreasedBy
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - contentLengthDecreasedBy cannot be negative
//
// contentLengthDecreasedBy:
//  - `2` represents len(wholeText)-2
//  - `0` represents start comparison from the end for both of the text.
func EndsWithPtr(
	wholeText, search *string,
	contentLengthDecreasedBy int,
	isCaseSensitive bool,
) bool {
	if wholeText == nil || search == nil {
		panic("Either wholeText or search is nil. Please provide valid string at least EmptyString(\"\")")
	}

	searchLength := len(*search)
	wholeTextLength := len(*wholeText)

	if searchLength == constants.Zero {
		return (wholeTextLength == constants.Zero && contentLengthDecreasedBy == constants.Zero) ||
			wholeTextLength-1 >= contentLengthDecreasedBy
	}

	if wholeTextLength == constants.Zero {
		return searchLength == constants.Zero && contentLengthDecreasedBy == constants.Zero
	}

	textLength := wholeTextLength - contentLengthDecreasedBy
	if searchLength > textLength {
		return false
	}

	if isCaseSensitive {
		return isstrinternal.EndsWith(
			wholeText,
			search,
			contentLengthDecreasedBy)
	}

	// insensitive
	return isEndsWithInsensitiveInternal(
		wholeText,
		search,
		contentLengthDecreasedBy)
}
