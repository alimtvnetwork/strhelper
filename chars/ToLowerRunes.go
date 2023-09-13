package chars

import "gitlab.com/auk-go/core/constants"

// ToLowerRunes
//
// Returns lower case runes by creating new runes. (Don't modify in place, thus requires more memory consumption)
//
// Invalid case (returns nil)
//   - if inputs == nil
func ToLowerRunes(inputs []rune) []rune {
	if inputs == nil {
		return nil
	}

	// Copying, example reference : https://play.golang.org/p/r65MrCg86YH
	toLowers := make([]rune, len(inputs))

	for index, r := range inputs {
		if r >= constants.UpperCaseA && r <= constants.UpperCaseZ {
			// in uppercase form, making it to lower case
			toLowers[index] = r + constants.LowerCase
		} else {
			toLowers[index] = r
		}
	}

	return toLowers
}
