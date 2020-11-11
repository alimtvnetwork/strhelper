package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Makes a new char to lower, doesn't modify the existing one.
func ToCharsUpper(chars *[]uint8) *[]uint8 {
	length := len(*chars)
	newChars := make([]uint8, length)

	for i := 0; i < length; i++ {
		char := (*chars)[i]

		if char >= constants.LowerCaseA &&
			char <= constants.LowerCaseZ {
			char = char + constants.UpperCaseA - constants.LowerCaseA
		}

		newChars[i] = char
	}

	return &newChars
}
