package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

// Returns the count number based on chars ([256]uint8 represents
// all ASCII characters  ASCII which index has flag 1) present in the str.
// Invalid Cases (return 0):
//  - str == nil or str == "" or length == 0
//
// @chars *[256]uint8:
//  - represents all ASCII characters in a simple array format,
//          only existing ones which are passed will be marked with 1.
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
		return CountAsciiCharsSensitivePtr(
			str,
			chars,
			startAt)
	}

	return CountAscIICharsInsensitivePtr(
		str,
		chars,
		startAt)
}
