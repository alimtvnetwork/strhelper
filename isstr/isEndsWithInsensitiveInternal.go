package isstr

import "strings"

// Results true for ends with search text. (case : Insensitive).
//
// Returns true
//  - if wholeText starts from the last with search text comparison.
//  - if contentLengthDecreasedBy mentioned then last len(wholeText)-contentLengthDecreasedBy
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - contentLengthDecreasedBy cannot be negative
//
// Warning:
//  - This doesn't do the quick exit, based on if search length > whole text length.
//      (Assumptions are it is already made before the call)
//
// contentLengthDecreasedBy:
//  - `2` represents len(wholeText)-2
//  - `0` represents start comparison from the end for both of the text.
func isEndsWithInsensitiveInternal(
	wholeText, search *string,
	contentLengthDecreasedBy int,
) bool {
	lenA := len(*wholeText)
	lenB := len(*search)

	wholeLower := strings.ToLower(*wholeText)
	searchLower := strings.ToLower(*search)

	incrementing := 0
	lastIndexWholeText := lenA - 1
	lastIndexSearchText := lenB - 1

	for ; contentLengthDecreasedBy < lenA && contentLengthDecreasedBy < lenB; contentLengthDecreasedBy++ {
		if (wholeLower)[lastIndexWholeText-contentLengthDecreasedBy] != (searchLower)[lastIndexSearchText-incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == lenB
}
