package strhelper

// startsAtLastIndex = 0 meaning starts from last position, giving 2 meaning len(wholeText)-2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
// For performance use Ptr version
func IsEndsWith(
	wholeText, endsWithSearch string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	return IsEndsWithPtr(
		&wholeText,
		&endsWithSearch,
		startsAtLastIndex,
		isCaseSensitive)
}
