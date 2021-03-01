package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

// Makes ascii to chars to lower case, modify existing chars.
//
// In terms of good practice, work with return value rather then existing one.
func ToCharsUpperInPlace(chars *[]uint8) *[]uint8 {
	for i, char := range *chars {
		if char >= constants.LowerCaseA &&
			char <= constants.LowerCaseZ {
			(*chars)[i] = char + constants.UpperCaseA - constants.LowerCaseA
		}
	}

	return chars
}
