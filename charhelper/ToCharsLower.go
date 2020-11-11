package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Makes a new char to lower, doesn't modify the existing one.
func ToCharsLower(chars *[]uint8) *[]uint8 {
	length := len(*chars)
	newChars := make([]uint8, length)

	for i := 0; i < length; i++ {
		char := (*chars)[i]

		if char >= constants.UpperCaseA &&
			char <= constants.UpperCaseZ {
			char = char + constants.LowerCase
		}

		newChars[i] = char
	}

	return &newChars
}
