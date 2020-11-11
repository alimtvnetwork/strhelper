package charhelper

// Counts and returns the count number based on chars present in the str
// Returns 0 if str is nil or empty string.
func CountAscIICharsSensitivePtr(
	str *string,
	chars *[256]uint8,
	at int,
) int {
	length := len(*str)

	if length == 0 {
		return 0
	}

	found := 0

	for ; at < length; at++ {
		char := (*str)[at]
		if chars[char] == 1 {
			found++
		}
	}

	return found
}
