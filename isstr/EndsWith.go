package isstr

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
func EndsWith(
	wholeText, endsWithSearch string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	return EndsWithPtr(
		&wholeText,
		&endsWithSearch,
		startsAtLastIndex,
		isCaseSensitive)
}
