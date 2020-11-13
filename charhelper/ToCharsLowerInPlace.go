package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Makes ascii to chars to upper case, modify existing chars.
//
// In terms of good practice, work with return value rather then existing one.
func ToCharsLowerInPlace(chars *[]uint8) *[]uint8 {
	for i, char := range *chars {
		if char >= constants.UpperCaseA &&
			char <= constants.UpperCaseZ {
			(*chars)[i] = char + constants.LowerCase
		}
	}

	return chars
}
