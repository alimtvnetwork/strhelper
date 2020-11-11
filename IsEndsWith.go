package strhelper

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
func IsEndsWith(
	wholeText, endsWithSearch string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	return IsEndsWithPtr(
		&wholeText,
		&endsWithSearch,
		startsAtLastIndex,
		isCaseSensitive)
}
