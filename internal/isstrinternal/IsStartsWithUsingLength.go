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
func IsStartsWithUsingLength(
	wholeText, search *string,
	startsAt int,
	wholeTextLength, searchLength int,
) bool {
	incrementing := 0

	for ; startsAt < wholeTextLength && incrementing < searchLength; startsAt++ {
		if (*wholeText)[startsAt] != (*search)[incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == searchLength
}
