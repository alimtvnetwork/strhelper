package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns Upper case runes by creating new runes. (Don't modify in place, thus requires more memory consumption)
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToUpperRunes(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	newRunes := make([]rune, len(*inputs))

	for index, rune := range *inputs {
		if rune >= constants.LowerCaseA && rune <= constants.LowerCaseZ {
			// in lower case form, making it to upper case
			rune = rune + constants.UpperCase
		}

		newRunes[index] = rune
	}

	return &newRunes
}
