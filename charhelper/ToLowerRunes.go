package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Creates new rune array of lowercase runes
// if inputs == nil then returns nil
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
