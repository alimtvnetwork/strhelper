package chars

import "gitlab.com/auk-go/core/constants"

// ToCharsUpper Makes a new char to Upper, doesn't modify the existing one.
//
// Warning: Requires double memory cost for copying the same data.
func ToCharsUpper(chars []uint8) []uint8 {
	// Copying, reference: https://play.golang.org/p/r65MrCg86YH
	toUppers := make([]uint8, len(chars))
	for i, char := range chars {
		if char >= constants.LowerCaseA &&
			char <= constants.LowerCaseZ {
			toUppers[i] = char + constants.UpperCaseA - constants.LowerCaseA
		} else {
			toUppers[i] = char
		}
	}

	return toUppers
}
