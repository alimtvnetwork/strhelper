package chars

import "gitlab.com/auk-go/core/constants"

// ToCharsLower Makes a new char to lower,
//
// doesn't modify the existing one.
func ToCharsLower(chars []uint8) []uint8 {
	// Copying, example reference : https://play.golang.org/p/r65MrCg86YH
	toLowers := make([]uint8, len(chars))
	for i, char := range chars {
		if char >= constants.UpperCaseA &&
			char <= constants.UpperCaseZ {
			toLowers[i] = char + constants.LowerCase
		} else {
			toLowers[i] = char
		}
	}

	return toLowers
}
