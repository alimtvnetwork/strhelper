package chars

import (
	"gitlab.com/auk-go/core/constants"
)

func ToLowerRune(r rune) rune {
	if r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ {
		lowerCaseRune := r + constants.LowerCase

		return lowerCaseRune
	}

	return r
}
