package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Makes ascii to chars to lower case, modify existing chars.
//
// In terms of good practice, work with return value rather then existing one.
func ToCharsUpperInPlace(chars *[]uint8) *[]uint8 {
	for i, char := range *chars {
		if char >= strconst.LowerCaseA &&
			char <= strconst.LowerCaseZ {
			(*chars)[i] = char + strconst.UpperCaseA - strconst.LowerCaseA
		}
	}

	return chars
}
