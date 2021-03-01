package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func IsUpperRuneCase(r rune) bool {
	return r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ
}
