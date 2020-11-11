package strhelper

func DoesntExist(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return !IsExists(s, findingString, isCaseSensitive)
}
