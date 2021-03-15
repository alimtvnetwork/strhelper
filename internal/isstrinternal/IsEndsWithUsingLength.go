package isstrinternal

// Results true for ends with search text. (case : Sensitive).
//
// Returns true
//  - if wholeText starts from the last with search text comparison.
//  - if lastIndexIncreasedBy mentioned then last len(wholeText)-lastIndexIncreasedBy
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - lastIndexIncreasedBy cannot be negative
//
// Warning:
//  - This doesn't do the quick exit, based on if search length > whole text length.
//      (Assumptions are it is already made before the call)
//
// lastIndexIncreasedBy:
//  - `2` represents len(wholeText)-2
//  - `0` represents start comparison from the end for both of the text.
func IsEndsWithUsingLength(
	wholeText, search *string,
	lastIndexIncreasedBy int,
	wholeTextLength,
	searchTextLength int,
) bool {
	incrementing := 0
	lastIndexWholeText := wholeTextLength - 1
	lastIndexSearchText := searchTextLength - 1

	for ; lastIndexIncreasedBy < wholeTextLength && lastIndexIncreasedBy < searchTextLength; lastIndexIncreasedBy++ {
		if (*wholeText)[lastIndexWholeText-lastIndexIncreasedBy] != (*search)[lastIndexSearchText-incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == searchTextLength
}
