package strhelper

// Returns all indexes where findingString is found.
// if empty content given then returns nil
// Invalid result can be nil if any (content == nil || findingString == nil) then returns nil
// for performance use Ptr version
func IndexesOfAll(
	content string,
	findingString string,
	startsAtIndex int,
	isCaseSensitive bool,
) []int {
	if content == findingString && startsAtIndex == 0 {
		return []int{0}
	}

	return IndexesOfAllPtr(
		&content,
		&findingString,
		startsAtIndex,
		isCaseSensitive)
}
