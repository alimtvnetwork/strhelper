package chars

import (
	"gitlab.com/auk-go/core/constants"
)

// ToCharsLowerInPlace Makes ascii to chars to upper case, modify existing chars.
//
// In terms of good practice, work with return value rather than existing one.
func ToCharsLowerInPlace(chars []uint8) []uint8 {
	for i, char := range chars {
		if char >= constants.UpperCaseA &&
			char <= constants.UpperCaseZ {
			chars[i] = char + constants.LowerCase
		}
	}

	return chars
}
