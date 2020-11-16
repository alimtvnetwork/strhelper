package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns lower case runes by modifying runes in place.
//
// Best practice is to work with the return value, even-though it was updated in place.
//
// Warning:
//  - inputs will be modified.
// Unhandled Case:
//  - if `inputs` is nil
func ToLowerRunesInPlace(inputs *[]rune) *[]rune {
	for index, r := range *inputs {
		if r >= strconst.UpperCaseA && r <= strconst.UpperCaseZ {
			// in uppercase form, making it to lower case
			(*inputs)[index] = r + strconst.LowerCase
		}
	}

	return inputs
}
