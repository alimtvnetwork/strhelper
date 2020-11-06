package strhelper

// startsAt = 0 meaning starts from last position, giving 2 meaning start index += 2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
// For performance use Ptr version
func IsStartsWith(
	wholeText, startsWith string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	return IsStartsWithPtr(
		&wholeText,
		&startsWith,
		startsAt,
		isCaseSensitive)
}
