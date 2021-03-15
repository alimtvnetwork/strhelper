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
func IsStartsWith(
	wholeText, search *string,
	startsAt int,
) bool {
	lenA := len(*wholeText)
	lenB := len(*search)
	incrementing := 0

	for ; startsAt < lenA && incrementing < lenB; startsAt++ {
		if (*wholeText)[startsAt] != (*search)[incrementing] {
			break
		}

		incrementing++
	}

	return incrementing == lenB
}
