package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
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
		if r >= strconst.LowerCaseA && r <= strconst.LowerCaseZ {
			// in lower case form, making it to upper case
			(*inputs)[index] = r + strconst.UpperCase
		}
	}

	return inputs
}
