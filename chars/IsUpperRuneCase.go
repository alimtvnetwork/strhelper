package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func IsUpperRuneCase(r rune) bool {
	return r >= strconst.UpperCaseA &&
		r <= strconst.UpperCaseZ
}
