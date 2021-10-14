package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

// Makes a new char to lower, doesn't modify the existing one.
//
// @chars *[256]uint8:
//  - represents all ASCII characters in a simple array format,
//      only existing ones which are passed will be marked with 1.
func ToAsciiCharsLower(chars [256]uint8) [256]uint8 {
	length := len(chars)
	newChars := [256]uint8{}

	for i := 0; i < length; i++ {
		if chars[i] == 1 {
			char := uint8(i)

			if char >= constants.UpperCaseA &&
				char <= constants.UpperCaseZ {
				char = char + constants.LowerCase
			}

			newChars[i] = char
		}
	}

	return newChars
}
