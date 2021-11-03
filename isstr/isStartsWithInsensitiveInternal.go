package isstr

import "strings"

// Results true for starts with (case : Insensitive).
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
//  - This doesn't do the quick exit, based on if search length > whole text length.
//      (Assumptions are it is already made before the call)
func isStartsWithInsensitiveInternal(
	wholeText,
	startsWith string,
	startsAt int,
	wholeTextLength, searchTextLength int,
) bool {
	wholeLower := strings.ToLower(wholeText)
	startsWithLower := strings.ToLower(startsWith)
	incrementing := 0

	for ; startsAt < wholeTextLength && incrementing < searchTextLength; startsAt++ {
		if wholeLower[startsAt] != startsWithLower[incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == searchTextLength
}
