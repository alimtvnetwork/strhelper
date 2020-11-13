package strhelper

// Returns all indexes found
//
// Returns : nil
//  - When not found or invalid case.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Conditions (for panic):
//  - startsAt cannot be negative or greater than the length of text(s)
//
// For performance use Ptr version
func IndexesOfAll(
	content string,
	findingString string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *[]int {
	if content == findingString && startsAtIndex == 0 {
		return &[]int{0}
	}

	request := createSearchRequest(
		&findingString,
		startsAtIndex,
		limits,
		isCaseSensitive,
	)

	return IndexesOfAllUsingRequestPtr(
		&content,
		request)
}
