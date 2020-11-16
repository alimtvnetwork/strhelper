package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func ToLower(c uint8) uint8 {
	if c >= strconst.UpperCaseA &&
		c <= strconst.UpperCaseZ {
		return c + strconst.LowerCase
	}

	return c
}
