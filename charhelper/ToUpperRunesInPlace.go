package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns Upper case runes by modifying runes in place.
//
// Best practice is to work with the return value, even-though it was updated in place.
//
// Warning:
//  - inputs will be modified.
// Invalid case (returns nil)
//  - if inputs == nil
func ToUpperRunesInPlace(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	for index, rune := range *inputs {
		if rune >= constants.LowerCaseA && rune <= constants.LowerCaseZ {
			// in lower case form, making it to upper case
			rune = rune + constants.UpperCase
			(*inputs)[index] = rune
		}
	}

	return inputs
}
