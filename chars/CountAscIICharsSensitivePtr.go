package chars

// Returns the count number based on chars ([256]uint8 represents all ASCII characters  ASCII which index has flag 1) present in the str. (Case : Sensitive)
// Invalid Cases (return 0):
//  - length == 0
//
// @chars *[256]uint8:
//  - represents all ASCII characters in a simple array format, only existing ones which are passed will be marked with 1.
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
