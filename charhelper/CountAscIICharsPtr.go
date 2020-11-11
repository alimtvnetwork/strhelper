package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Counts and returns the count number based on chars present in the str
// Returns 0 if str is nil or empty string.
func CountAscIICharsPtr(
	str *string,
	chars *[256]uint8,
	startAt int,
	isCaseSensitive bool,
) int {
	if str == nil || *str == constants.EmptyString || len(*str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountAscIICharsSensitivePtr(
			str,
			chars,
			startAt)
	}

	return CountAscIICharsInsensitivePtr(
		str,
		chars,
		startAt)
}
