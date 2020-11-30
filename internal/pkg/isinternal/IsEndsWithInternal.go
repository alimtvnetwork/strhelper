package isinternal

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
func IsEndsWithInternal(
	wholeText, search *string,
	lastIndexIncreasedBy int,
) bool {
	lenA := len(*wholeText) - lastIndexIncreasedBy
	lenB := len(*search)
	lastIndexIncreasedBy = 0

	for ; lastIndexIncreasedBy < lenA && lastIndexIncreasedBy < lenB; lastIndexIncreasedBy++ {
		if (*wholeText)[lenA-1-lastIndexIncreasedBy] != (*search)[lenB-1-lastIndexIncreasedBy] {
			break
		}
	}

	return lastIndexIncreasedBy == lenB
}
