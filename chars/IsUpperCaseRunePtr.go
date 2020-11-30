package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func IsUpperCaseRunePtr(r *rune) bool {
	return *r >= strconst.UpperCaseA &&
		*r <= strconst.UpperCaseZ
}
