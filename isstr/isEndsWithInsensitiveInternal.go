package isstr

import "strings"

// Results true for ends with search text. (case : Insensitive).
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
func isEndsWithInsensitiveInternal(
	wholeText, search *string,
	lastIndexIncreasedBy int,
) bool {
	lenA := len(*wholeText) - lastIndexIncreasedBy
	lenB := len(*search)

	wholeLower := strings.ToLower(*wholeText)
	searchLower := strings.ToLower(*search)
	lastIndexIncreasedBy = 0

	for ; lastIndexIncreasedBy < lenA && lastIndexIncreasedBy < lenB; lastIndexIncreasedBy++ {
		if wholeLower[lenA-1-lastIndexIncreasedBy] != searchLower[lenB-1-lastIndexIncreasedBy] {
			break
		}
	}

	return lastIndexIncreasedBy == lenB
}
