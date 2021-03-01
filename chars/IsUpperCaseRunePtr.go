package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func IsUpperCaseRunePtr(r *rune) bool {
	return *r >= constants.UpperCaseA &&
		*r <= constants.UpperCaseZ
}
