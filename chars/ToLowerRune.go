package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func ToLowerRune(r rune) rune {
	if r >= strconst.UpperCaseA &&
		r <= strconst.UpperCaseZ {
		lowerCaseRune := r + strconst.LowerCase

		return lowerCaseRune
	}

	return r
}
