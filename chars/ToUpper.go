package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func ToUpper(c uint8) uint8 {
	if c >= strconst.LowerCaseA &&
		c <= strconst.LowerCaseZ {
		return c + strconst.UpperCaseA - strconst.LowerCaseA
	}

	return c
}
