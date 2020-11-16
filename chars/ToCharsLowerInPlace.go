package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Makes ascii to chars to upper case, modify existing chars.
//
// In terms of good practice, work with return value rather then existing one.
func ToCharsLowerInPlace(chars *[]uint8) *[]uint8 {
	for i, char := range *chars {
		if char >= strconst.UpperCaseA &&
			char <= strconst.UpperCaseZ {
			(*chars)[i] = char + strconst.LowerCase
		}
	}

	return chars
}
