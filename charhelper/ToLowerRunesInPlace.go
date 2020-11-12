package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns lower case runes by modifying runes in place.
// if inputs == nil then returns nil
//
// Best practice is to work with the return value, even-though it was updated in place.
//
// Warning:
//  - inputs will be modified.
func ToLowerRunesInPlace(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	for index, rune := range *inputs {
		if rune >= constants.UpperCaseA && rune <= constants.UpperCaseZ {
			// in uppercase form, making it to lower case
			rune = rune + constants.LowerCase
			(*inputs)[index] = rune
		}
	}

	return inputs
}
