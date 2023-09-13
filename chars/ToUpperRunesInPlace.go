package chars

import (
	"gitlab.com/auk-go/core/constants"
)

// Returns Upper case runes by modifying runes in place.
//
// Best practice is to work with the return value, even-though it was updated in place.
//
// Warning:
//  - inputs will be modified.
// Unhandled Case:
//  - if `inputs` is nil
func ToUpperRunesInPlace(inputs *[]rune) *[]rune {
	for index, r := range *inputs {
		if r >= constants.LowerCaseA && r <= constants.LowerCaseZ {
			// in lower case form, making it to upper case
			(*inputs)[index] = r + constants.UpperCase
		}
	}

	return inputs
}
