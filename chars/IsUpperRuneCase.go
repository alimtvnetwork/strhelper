package chars

import (
	"gitlab.com/auk-go/core/constants"
)

func IsUpperRuneCase(r rune) bool {
	return r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ
}
