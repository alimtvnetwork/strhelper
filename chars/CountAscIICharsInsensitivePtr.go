package chars

import "strings"

// Counts and returns the count number based on chars present in the str
// Returns 0 if str is nil or empty string.
func CountAscIICharsInsensitivePtr(
	str *string,
	chars *[256]uint8,
	at int,
) int {
	length := len(*str)
	found := 0
	strLower := strings.ToLower(*str)
	charsLower := ToAsciiCharsLower(chars)

	for ; at < length; at++ {
		char := strLower[at]
		if charsLower[char] == 1 {
			found++
		}
	}

	return found
}
