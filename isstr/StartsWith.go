package isstr

// Results true for starts with.
//
// Returns true
//
//  - if wholeText starts with search text from the index mentioned at startsAt.
//
// Conditions (Not Handled and Assumptions):
//  - wholeText, search should NOT be nil.
//  - startsAt cannot be negative
//
// For better performance use `...Ptr` version of the method.
func StartsWith(
	wholeText, startsWith string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	return StartsWithPtr(
		&wholeText,
		&startsWith,
		startsAt,
		isCaseSensitive)
}
