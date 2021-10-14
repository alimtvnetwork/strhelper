package chars

// Counts and returns the count number based on chars present in the str
// Returns 0 if str is nil or empty string.
func CountAscIIChars(
	str string,
	chars [256]uint8,
	startsAt int,
	isCaseSensitive bool,
) int {
	return CountAscIICharsPtr(
		str,
		chars,
		startsAt,
		isCaseSensitive)
}
