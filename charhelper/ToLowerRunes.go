package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns lower case runes by creating new runes. (Don't modify in place, thus requires more memory consumption)
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToLowerRunes(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	newRunes := make([]rune, len(*inputs))

	for index, rune := range *inputs {
		if rune >= constants.UpperCaseA && rune <= constants.UpperCaseZ {
			// in uppercase form, making it to lower case
			rune = rune + constants.LowerCase
		}

		newRunes[index] = rune
	}

	return &newRunes
}
