package isstrinternal

// Results true for starts with (case : Sensitive).
//
// Returns true
//
//  - if wholeText starts with search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
//
// Warning:
// - This doesn't do the quick exit, based on if search length > whole text length.
// (Assumptions are it is already made before the call)
func StartsWith(
	wholeText, search *string,
	startsAt int,
	wholeTextLength, searchTextLength int,
) bool {
	incrementing := 0
	// accessing direct without pointer increases performance
	wholeTextCopy := *wholeText
	searchTextCopy := *search

	for ; startsAt < wholeTextLength && incrementing < searchTextLength; startsAt++ {
		if wholeTextCopy[startsAt] != searchTextCopy[incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == searchTextLength
}
